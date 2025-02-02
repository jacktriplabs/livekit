// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package streamallocator

import (
	"fmt"
	"time"

	"github.com/livekit/protocol/logger"
)

// ------------------------------------------------

type NackTrackerConfig struct {
	WindowMinDuration time.Duration `yaml:"window_min_duration,omitempty"`
	WindowMaxDuration time.Duration `yaml:"window_max_duration,omitempty"`
	RatioThreshold    float64       `yaml:"ratio_threshold,omitempty"`
}

var (
	DefaultNackTrackerConfigProbe = NackTrackerConfig{
		WindowMinDuration: 500 * time.Millisecond,
		WindowMaxDuration: 1 * time.Second,
		RatioThreshold:    0.04,
	}

	DefaultNackTrackerConfigNonProbe = NackTrackerConfig{
		WindowMinDuration: 2 * time.Second,
		WindowMaxDuration: 3 * time.Second,
		RatioThreshold:    0.08,
	}
)

// ------------------------------------------------

type NackTrackerParams struct {
	Name   string
	Logger logger.Logger
	Config NackTrackerConfig
}

type NackTracker struct {
	params NackTrackerParams

	windowStartTime time.Time
	packets         uint32
	repeatedNacks   uint32

	/* STREAM-ALLOCATOR-DATA
	// STREAM-ALLOCATOR-EXPERIMENTAL-TODO: remove when cleaning up experimental stuff
	history []string
	*/
}

func NewNackTracker(params NackTrackerParams) *NackTracker {
	return &NackTracker{
		params: params,
		// STREAM-ALLOCATOR-DATA history: make([]string, 0, 10),
	}
}

func (n *NackTracker) Add(packets uint32, repeatedNacks uint32) {
	if n.params.Config.WindowMaxDuration != 0 && !n.windowStartTime.IsZero() && time.Since(n.windowStartTime) > n.params.Config.WindowMaxDuration {
		// STREAM-ALLOCATOR-DATA n.updateHistory()

		n.windowStartTime = time.Time{}
		n.packets = 0
		n.repeatedNacks = 0
	}

	//
	// Start NACK monitoring window only when a repeated NACK happens.
	// This allows locking tightly to when NACKs start happening and
	// check if the NACKs keep adding up (potentially a sign of congestion)
	// or isolated losses
	//
	if n.repeatedNacks == 0 && repeatedNacks != 0 {
		n.windowStartTime = time.Now()
	}

	if !n.windowStartTime.IsZero() {
		n.packets += packets
		n.repeatedNacks += repeatedNacks
	}
}

func (n *NackTracker) GetRatio() float64 {
	ratio := 0.0
	if n.packets != 0 {
		ratio = float64(n.repeatedNacks) / float64(n.packets)
		if ratio > 1.0 {
			ratio = 1.0
		}
	}

	return ratio
}

func (n *NackTracker) IsTriggered() bool {
	if n.params.Config.WindowMinDuration != 0 && !n.windowStartTime.IsZero() && time.Since(n.windowStartTime) > n.params.Config.WindowMinDuration {
		return n.GetRatio() > n.params.Config.RatioThreshold
	}

	return false
}

func (n *NackTracker) ToString() string {
	window := ""
	if !n.windowStartTime.IsZero() {
		now := time.Now()
		elapsed := now.Sub(n.windowStartTime).Seconds()
		window = fmt.Sprintf("t: %+v|%+v|%.2fs", n.windowStartTime.Format(time.UnixDate), now.Format(time.UnixDate), elapsed)
	}
	return fmt.Sprintf("n: %s, %s, p: %d, rn: %d, rn/p: %.2f", n.params.Name, window, n.packets, n.repeatedNacks, n.GetRatio())
}

/* STREAM-ALLOCATOR-DATA
func (n *NackTracker) GetHistory() []string {
	return n.history
}

func (n *NackTracker) updateHistory() {
	if len(n.history) >= 10 {
		n.history = n.history[1:]
	}

	n.history = append(n.history, n.ToString())
}
*/

// ------------------------------------------------
