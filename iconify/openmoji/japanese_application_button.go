package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseApplicationButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07a.968.968 0 0 1 .965-.965h46.07a.968.968 0 0 1 .965.965v46.07a.968.968 0 0 1-.965.965Z"/><g stroke="#000" stroke-width="2"><path fill="none" stroke-miterlimit="10" d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07a.968.968 0 0 1 .965-.965h46.07a.968.968 0 0 1 .965.965v46.07a.968.968 0 0 1-.965.965Z"/><path fill="none" stroke-linecap="round" stroke-linejoin="round" d="M25.5 41.5h21m0-8h-21m22 10v-17h-23v17"/><path stroke-linecap="round" stroke-linejoin="round" d="M36.5 21.5v30"/></g>`),
		g.Group(children),
	)
}