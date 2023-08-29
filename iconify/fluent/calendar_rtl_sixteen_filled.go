package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarRtlSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 11.5V6h12v5.5a2.5 2.5 0 0 1-2.5 2.5h-7A2.5 2.5 0 0 1 2 11.5Zm8.003-3.252a.748.748 0 1 0 1.497 0a.748.748 0 0 0-1.497 0Zm0 2.5a.748.748 0 1 0 1.497 0a.748.748 0 0 0-1.497 0Zm-2.752-2.5a.748.748 0 1 0 1.497 0a.748.748 0 0 0-1.497 0Zm0 2.5a.748.748 0 1 0 1.497 0a.748.748 0 0 0-1.497 0Zm-2.748-2.5a.748.748 0 1 0 1.497 0a.748.748 0 0 0-1.497 0ZM2 4.5A2.5 2.5 0 0 1 4.5 2h7A2.5 2.5 0 0 1 14 4.5V5H2v-.5Z"/>`),
		g.Group(children),
	)
}