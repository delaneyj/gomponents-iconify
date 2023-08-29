package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M320 291.24V80a64 64 0 1 0-128 0v211.24A113.39 113.39 0 0 0 144 384a112 112 0 0 0 224 0a113.39 113.39 0 0 0-48-92.76ZM256 432a48 48 0 0 1-16-93.26V96h32v242.74A48 48 0 0 1 256 432Z"/>`),
		g.Group(children),
	)
}