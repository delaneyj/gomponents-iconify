package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RamenDining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22v-1.75Q5.325 19.2 3.662 17T2 12h2V4l18-2v1.5L10.5 4.8v1.7H22V8H10.5v4H22q0 2.8-1.663 5T16 20.25V22H8ZM8 6.5h1V4.95l-1 .1V6.5Zm-2.5 0h1V5.25l-1 .1V6.5ZM8 12h1V8H8v4Zm-2.5 0h1V8h-1v4Z"/>`),
		g.Group(children),
	)
}