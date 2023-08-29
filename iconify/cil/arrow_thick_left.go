package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m294.637 496l-38.688-.035L16 255.729L256.334 16.048h38.277l.008 143.937h200.017v192H294.629ZM61.271 255.773l201.364 201.6l-.008-137.391h200.009v-128H262.621l-.008-137.006Z"/>`),
		g.Group(children),
	)
}