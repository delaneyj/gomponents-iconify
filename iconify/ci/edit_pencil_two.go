package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditPencilTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v4h4L18.869 9.131h0c.396-.396.594-.594.668-.822a1 1 0 0 0 0-.618c-.074-.228-.272-.426-.668-.822l-1.74-1.74c-.395-.394-.592-.592-.82-.666a1 1 0 0 0-.618 0c-.228.074-.426.272-.82.667l-.002.001L4 16.001Z"/>`),
		g.Group(children),
	)
}