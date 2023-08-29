package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopwatchStart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTStopwatchStart0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M24 44c9.389 0 17-7.611 17-17s-7.611-17-17-17S7 17.611 7 27s7.611 17 17 17Z"/><path stroke-linecap="round" d="M18 4h12m-6 15v8m8 0h-8m0-23v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTStopwatchStart0)"/>`),
		g.Group(children),
	)
}