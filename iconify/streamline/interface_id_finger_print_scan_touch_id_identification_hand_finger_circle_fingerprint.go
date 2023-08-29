package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceIdFingerPrintScanTouchIdIdentificationHandFingerCircleFingerprint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 13.5V8.25A1.25 1.25 0 0 1 6.75 7h0A1.25 1.25 0 0 1 8 8.25V11h2a2 2 0 0 1 2 2v.5"/><path d="M3.39 8.61a4.75 4.75 0 1 1 6.72 0"/></g>`),
		g.Group(children),
	)
}