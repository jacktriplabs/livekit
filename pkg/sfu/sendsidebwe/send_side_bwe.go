package sendsidebwe

import (
	"fmt"

	"github.com/livekit/protocol/logger"
)

//
// Based on a simplified/modified version of JitterPath paper
// (https://homepage.iis.sinica.edu.tw/papers/lcs/2114-F.pdf)
//
// TWCC feedback is uesed to calcualte delta one-way-delay.
// It is accumulated/propagated to determine in which region
// groups of packets are operating in.
//
// In simplified terms,
//   o JQR (Join Queuing Region) is when channel is congested.
//   o DQR (Disjoint Queuing Region) is when channel is not.
//
// Packets are grouped and thresholds applied to smooth over
// small variations. For example, in the paper,
//    if propagated_queuing_delay + delta_one_way_delay > 0 {
//       possibly_operating_in_jqr
//    }
// But, in this implementation it is checked at packet group level,
// i. e. using queuing delay and aggreated delta one-way-delay of
// the group and a minimum value threshold is applied before declaring
// that a group is in JQR.
//
// There is also hysteresis to make transisitons smoother, i.e. if the
// metric is above a certain threshold, it is JQR and it is DQR only if it
// is below a certain value and the gap in between those two thresholds
// are treated as interdeterminate groups.
//

// ---------------------------------------------------------------------------

type CongestionState int

const (
	CongestionStateNone CongestionState = iota
	CongestionStateEarlyWarning
	CongestionStateEarlyWarningHangover
	CongestionStateCongested
	CongestionStateCongestedHangover
)

func (c CongestionState) String() string {
	switch c {
	case CongestionStateNone:
		return "NONE"
	case CongestionStateEarlyWarning:
		return "EARLY_WARNING"
	case CongestionStateEarlyWarningHangover:
		return "EARLY_WARNING_HANGOVER"
	case CongestionStateCongested:
		return "CONGESTED"
	case CongestionStateCongestedHangover:
		return "CONGESTED_HANGOVER"
	default:
		return fmt.Sprintf("%d", int(c))
	}
}

// ---------------------------------------------------------------------------

type SendSideBWEConfig struct {
	CongestionDetector CongestionDetectorConfig `yaml:"congestion_detector,omitempty"`
}

var (
	DefaultSendSideBWEConfig = SendSideBWEConfig{
		CongestionDetector: DefaultCongestionDetectorConfig,
	}
)

// ---------------------------------------------------------------------------

type SendSideBWEParams struct {
	Config SendSideBWEConfig
	Logger logger.Logger
}

type SendSideBWE struct {
	params SendSideBWEParams

	*congestionDetector
}

func NewSendSideBWE(params SendSideBWEParams) *SendSideBWE {
	return &SendSideBWE{
		params: params,
		congestionDetector: newCongestionDetector(congestionDetectorParams{
			Config: params.Config.CongestionDetector,
			Logger: params.Logger,
		}),
	}
}

func (s *SendSideBWE) Stop() {
	s.congestionDetector.Stop()
}

// ------------------------------------------------
