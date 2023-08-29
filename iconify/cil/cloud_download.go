package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M272 434.744V209.176h-32v225.568l-51.882-51.882l-22.628 22.627L256 496l90.51-90.511l-22.628-22.627L272 434.744z"/><path fill="currentColor" d="M400 161.176c0-79.4-64.6-144-144-144s-144 64.6-144 144a96 96 0 0 0 0 192h80v-32h-80a64 64 0 0 1 0-128h32v-32a112 112 0 0 1 224 0v32h32a64 64 0 0 1 0 128h-80v32h80a96 96 0 0 0 0-192Z"/>`),
		g.Group(children),
	)
}