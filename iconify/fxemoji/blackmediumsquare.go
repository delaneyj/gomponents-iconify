package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blackmediumsquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#132028" d="M465.833 70.79c0-13.543-11.081-24.624-24.624-24.624H67.995c-12.189 0-22.161 9.973-22.161 22.161v375.677c0 12.189 9.973 22.161 22.161 22.161H441.21c13.543 0 24.624-11.081 24.624-24.624V70.79z"/>`),
		g.Group(children),
	)
}