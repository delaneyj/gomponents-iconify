package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareFSevenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M18.333 2c1.96 0 3.56 1.537 3.662 3.472l.005.195v12.666c0 1.96-1.537 3.56-3.472 3.662l-.195.005H5.667a3.667 3.667 0 0 1-3.662-3.472L2 18.333V5.667c0-1.96 1.537-3.56 3.472-3.662L5.667 2h12.666zM16 8h-3l-.117.007A1 1 0 0 0 12 9l.007.117A1 1 0 0 0 13 10h1.718l-1.188 4.757l-.022.115a1 1 0 0 0 1.962.37l1.5-6l.021-.11A1 1 0 0 0 16 8zm-6 0H8l-.117.007A1 1 0 0 0 7 9v6l.007.117A1 1 0 0 0 8 16l.117-.007A1 1 0 0 0 9 15v-2h1l.117-.007A1 1 0 0 0 10 11H9v-1h1l.117-.007A1 1 0 0 0 10 8z"/></g>`),
		g.Group(children),
	)
}