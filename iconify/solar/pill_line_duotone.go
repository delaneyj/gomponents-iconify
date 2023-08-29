package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.99 13.602a6.796 6.796 0 0 1 9.612-9.611l6.407 6.407a6.796 6.796 0 1 1-9.61 9.611L3.99 13.602Z"/><path d="M16.806 7.194s-.541 2.806-3.674 5.939C10 16.265 7.194 16.806 7.194 16.806" opacity=".5"/></g>`),
		g.Group(children),
	)
}