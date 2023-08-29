package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func File(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 512h343q13 0 22.5-9.5T407 480V32q-2-13-11.5-22.5T373 0H183L0 192v288q0 13 9.5 22.5T32 512zM192 43v128q0 13-4 17t-17 4H43zM43 235h128q64 0 64-64V43h128v426H43V235z"/>`),
		g.Group(children),
	)
}