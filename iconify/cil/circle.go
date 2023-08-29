package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256.6 496A239.364 239.364 0 0 0 425.856 87.379A239.364 239.364 0 0 0 87.344 425.892A237.8 237.8 0 0 0 256.6 496Zm0-446.729c114.341 0 207.365 93.023 207.365 207.364S370.941 464 256.6 464S49.236 370.977 49.236 256.635S142.259 49.271 256.6 49.271Z"/>`),
		g.Group(children),
	)
}