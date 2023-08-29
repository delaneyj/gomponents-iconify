package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fsharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 12L11.39.61v5.695L5.695 12l5.695 5.695v5.695L0 12zm7.322 0l4.068-4.068v8.136L7.322 12zM24 12L12.203.61v5.695L17.898 12l-5.695 5.695v5.695L24 12z"/>`),
		g.Group(children),
	)
}