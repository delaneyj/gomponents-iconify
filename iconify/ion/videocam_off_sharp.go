package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideocamOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m57.376 79.998l22.627-22.627L438.633 416l-22.627 22.628zM32 112a16 16 0 0 0-16 16v256a16 16 0 0 0 16 16h288a15.89 15.89 0 0 0 9.34-3l-285-285Zm304 96v-80a16 16 0 0 0-16-16H179.63l245.44 245.44L496 400V112Z"/>`),
		g.Group(children),
	)
}