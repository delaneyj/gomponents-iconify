package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrushAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M56 256v240h400V256a48.055 48.055 0 0 0-48-48h-86.617l13.075-106.263l.068-.677a78.777 78.777 0 1 0-157.052 0l.027.338L190.617 208H104a48.055 48.055 0 0 0-48 48Zm368 208H88V320h336ZM226.8 240L209.348 98.192a46.777 46.777 0 1 1 93.3 0L285.205 240H408a16.019 16.019 0 0 1 16 16v32H88v-32a16.019 16.019 0 0 1 16-16Z"/>`),
		g.Group(children),
	)
}