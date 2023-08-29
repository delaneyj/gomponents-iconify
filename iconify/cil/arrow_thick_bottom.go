package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M255.682 494.636L16 254.3v-38.276l143.937-.007V16h192v200.007L495.952 216l-.035 38.688ZM54.931 248.022l200.8 201.342L457.328 248l-137.391.008V48h-128v200.015Z"/>`),
		g.Group(children),
	)
}