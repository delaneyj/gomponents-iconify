package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Save(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSave0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M6 9a3 3 0 0 1 3-3h25.281L42 13.207V39a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9Z"/><path fill="#000" fill-rule="evenodd" d="M24.008 6L24 13.385c0 .34-.448.615-1 .615h-8c-.552 0-1-.275-1-.615V6" clip-rule="evenodd"/><path stroke="#000" stroke-linejoin="round" stroke-width="4" d="M24.008 6L24 13.385c0 .34-.448.615-1 .615h-8c-.552 0-1-.275-1-.615V6h10.008Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M9 6h25.281"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 26h20m-20 8h10.008"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSave0)"/>`),
		g.Group(children),
	)
}