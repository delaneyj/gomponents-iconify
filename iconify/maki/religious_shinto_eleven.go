package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReligiousShintoEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9.25 2.009h-7.5a.75.75 0 0 0 0 1.5H2v5.25a.75.75 0 0 0 1.5 0v-2.75h4v2.75a.75.75 0 0 0 1.5 0v-5.25h.25a.75.75 0 0 0 0-1.5zM7.5 4.759H6v-1.25h1.5zm-4-1.25H5v1.25H3.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}