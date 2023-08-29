package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icecream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M608 320H32q-13 0-22.5-9.5T0 288q0-34 15-53.5T64 192q3-78 59.5-135T256 0q73 0 127 49t63 121q23-10 50-10q61 0 102.5 35t41.5 93q0 13-9.5 22.5T608 320zM64 384h512q26 0 45 19t19 45L384 960q0 26-19 45t-45.5 19t-45-19t-18.5-45L0 448q0-26 19-45t45-19z"/>`),
		g.Group(children),
	)
}