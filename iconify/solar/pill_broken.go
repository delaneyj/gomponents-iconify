package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M2.6 6a6.796 6.796 0 0 1 11.002-2.01l6.407 6.408A6.765 6.765 0 0 1 21.893 14M2.107 10a6.765 6.765 0 0 0 1.884 3.602l6.407 6.407A6.796 6.796 0 0 0 21.4 18"/><path d="M16.806 7.194s-.541 2.806-3.674 5.939C10 16.265 7.194 16.806 7.194 16.806"/></g>`),
		g.Group(children),
	)
}