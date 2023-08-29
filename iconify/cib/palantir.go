package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Palantir(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.865 24L16 28.24L5.135 24l-1.802 3.125L16 32l12.667-4.875zM16 0C9.005 0 3.333 5.672 3.333 12.667S9.005 25.334 16 25.334s12.667-5.672 12.667-12.667S22.995 0 16 0zm0 21.438a8.754 8.754 0 0 1-8.755-8.755c0-4.839 3.917-8.76 8.755-8.76s8.755 3.922 8.755 8.76A8.755 8.755 0 0 1 16 21.438z"/>`),
		g.Group(children),
	)
}