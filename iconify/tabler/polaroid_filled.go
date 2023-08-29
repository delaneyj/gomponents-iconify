package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PolaroidFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="m9.199 9.623l.108.098l3.986 3.986l.094.083a1 1 0 0 0 1.403-1.403l-.083-.094l-.292-.293l1.292-1.293l.106-.095c.457-.38.918-.38 1.386.011l.108.098l4.502 4.503a4.003 4.003 0 0 1-3.596 2.77L18 18H6a4.002 4.002 0 0 1-3.809-2.775l5.516-5.518l.106-.095c.457-.38.918-.38 1.386.011zM18 2a4 4 0 0 1 3.995 3.8L22 6v6.585l-3.293-3.292l-.15-.137c-1.256-1.095-2.85-1.097-4.096-.017l-.154.14L13 10.585l-2.293-2.292l-.15-.137c-1.256-1.095-2.85-1.097-4.096-.017l-.154.14L2 12.585V6a4 4 0 0 1 3.8-3.995L6 2h12zm-2.99 3l-.127.007a1 1 0 0 0 0 1.986L15 7l.127-.007a1 1 0 0 0 0-1.986L15.01 5zm-7 15a1 1 0 0 1 .117 1.993L8 22a1 1 0 0 1-.117-1.993L8.01 20zm4 0a1 1 0 0 1 .117 1.993L12 22a1 1 0 0 1-.117-1.993L12.01 20zm4 0a1 1 0 0 1 .117 1.993L16 22a1 1 0 0 1-.117-1.993L16.01 20z"/></g>`),
		g.Group(children),
	)
}