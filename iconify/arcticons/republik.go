package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Republik(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.208 23.93c11.544-.736 12.257-7.126 12.257-9.257c0-2.262-1.095-8.319-8.393-8.319H5.5l3.649 2.773v29.556L5.5 41.457l17.807.19l-3.65-2.774V27.737l3.081-2.548l6.99 16.457H42.5l-2.846-2.744l-6.568-12.552l-7.856-1.82l-.022-.6Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.87 22.02c3.017 0 6.691-1.906 6.691-6.476c0-3.23-.888-5.72-6.075-7.15l-3.202 2.847v10.78h2.585Z"/>`),
		g.Group(children),
	)
}