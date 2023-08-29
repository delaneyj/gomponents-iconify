package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path fill="currentColor" d="M12 2H2a2 2 0 0 0-2 2v8c0 5.982 6 11 12 11s12-5.018 12-11V4a2 2 0 0 0-2-2H12Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="m6 9l6.404 6L18 9"/></g>`),
		g.Group(children),
	)
}