package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderMale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m12.189 7l.002-2l7 .007l-.008 7l-2-.002l.004-3.588l-3.044 3.044a5.002 5.002 0 0 1-7.679 6.336a5 5 0 0 1 6.25-7.736l3.058-3.057L12.189 7Zm-4.31 5.14a3 3 0 1 1 4.242 4.244A3 3 0 0 1 7.88 12.14Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}