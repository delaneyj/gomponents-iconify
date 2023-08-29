package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microsoft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22h9.5v-9.5H2V22zm0-10.5h9.5V2H2v9.5zM12.5 2v9.5H22V2h-9.5zm0 20H22v-9.5h-9.5V22z"/>`),
		g.Group(children),
	)
}