package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceShareMegaPhoneOneBullhornLoudMegaphoneShareSpeakerTransmit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.84 12.07L1.17 8a1 1 0 0 1-.67-.91V6a1 1 0 0 1 .67-.94L12.84 1a.5.5 0 0 1 .66.47V11.6a.5.5 0 0 1-.66.47Zm-4.36-1.5A2.75 2.75 0 0 1 3 10.25V8.66"/>`),
		g.Group(children),
	)
}