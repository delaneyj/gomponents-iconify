package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MirrorLeftBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M3 12c0 4.243 0 6.364 1.318 7.682C5.636 21 7.758 21 12 21m0-18C7.757 3 5.636 3 4.318 4.318C3.502 5.134 3.192 6.257 3.073 8"/><path stroke-dasharray="2.5 3" d="M11 21h4c2.828 0 4.243 0 5.121-.879C21 19.243 21 17.828 21 15V9c0-2.828 0-4.243-.879-5.121C19.243 3 17.828 3 15 3h-4"/><path d="M12 22V2"/></g>`),
		g.Group(children),
	)
}