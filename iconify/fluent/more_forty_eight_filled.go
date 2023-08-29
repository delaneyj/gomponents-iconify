package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M16.668 24a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0z" fill="currentColor"/><path d="M28.168 24a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0z" fill="currentColor"/><path d="M36.168 27.5a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}