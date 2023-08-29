package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12.3L3.5 7.05L12 1.8l8.5 5.25z"/><path fill="currentColor" d="M12 22.2v-9.9l8.5-5.25v9.9z" opacity=".25"/><path fill="currentColor" d="m12 22.2l-8.5-5.25v-9.9L12 12.3z" opacity=".5"/>`),
		g.Group(children),
	)
}