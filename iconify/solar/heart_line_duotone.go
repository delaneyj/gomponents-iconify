package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M12 5.5C7.5.826 2 4.275 2 9.138C2 14 6.02 16.591 8.962 18.911C10 19.729 11 20.5 12 20.5"/><path d="M12 5.5C16.5.826 22 4.275 22 9.138c0 4.863-4.02 7.454-6.962 9.774C14 19.729 13 20.5 12 20.5" opacity=".5"/></g>`),
		g.Group(children),
	)
}