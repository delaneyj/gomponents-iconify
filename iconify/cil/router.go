package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Router(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m443.057 132.943l22.634-22.634a143.764 143.764 0 0 0-211.382 0l22.634 22.634a111.838 111.838 0 0 1 166.114 0Z"/><path fill="currentColor" d="m299.615 155.615l22.7 22.7a47.913 47.913 0 0 1 75.362 0l22.7-22.7a79.829 79.829 0 0 0-120.77 0ZM472 280h-96v-64h-32v64H40a24.028 24.028 0 0 0-24 24v112a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V304a24.028 24.028 0 0 0-24-24Zm-8 128H48v-96h416Z"/><path fill="currentColor" d="M96 344h32v32H96zm80 0h32v32h-32zm80 0h32v32h-32z"/>`),
		g.Group(children),
	)
}