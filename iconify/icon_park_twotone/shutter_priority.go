package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShutterPriority(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTShutterPriority0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="m15 12l3-6h12l3 6H15Z"/><path fill="#555" d="M41 12H7c-1.657 0-3 1.254-3 2.8v24.4C4 40.746 5.343 42 7 42h34c1.657 0 3-1.254 3-2.8V14.8c0-1.546-1.343-2.8-3-2.8Z"/><path stroke-linecap="round" d="M28 20c-3 .13-9 1.089-9 3.889c0 3.5 10 2.722 10 6.222c0 2.8-6.667 3.76-10 3.889"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTShutterPriority0)"/>`),
		g.Group(children),
	)
}