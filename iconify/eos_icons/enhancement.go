package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Enhancement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.02 13.99h4v8.02h-4zm7-3.97h4v11.99h-4zM22 6l-4-4l-4 4h2v16.02h4V6h2z"/>`),
		g.Group(children),
	)
}