package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spikedshoes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSpikedshoes0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path stroke="#fff" d="M44 6H28v8h16V6Z"/><path fill="#fff" stroke="#fff" d="M44 14v22c0 1.11-.89 2-2 2H8c-2.21 0-4-1.79-4-4v-6c0-4.42 3.58-8 8-8h16v-6h16Z"/><path stroke="#000" d="M14 26v-6m7 6v-6"/><path stroke="#fff" d="M15 42v-4m-7 4v-4m14 4v-4m12 4v-4m7 4v-4M23 20H12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSpikedshoes0)"/>`),
		g.Group(children),
	)
}