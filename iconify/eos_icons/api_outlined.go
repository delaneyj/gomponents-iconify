package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApiOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.312 9H5.688L3.5 15h1.607l.446-1.226h1.894L7.893 15H9.5Zm-1.394 3.774L6.5 11.18l.582 1.595ZM14.744 9h-3.5v6h1.5v-2h2a1.473 1.473 0 0 0 1.5-1.5v-1a1.473 1.473 0 0 0-1.5-1.5Zm0 2.5h-2v-1h2ZM18 9h1.5v6H18z"/><path fill="currentColor" d="M22 6v12H2V6h20m0-2H2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}