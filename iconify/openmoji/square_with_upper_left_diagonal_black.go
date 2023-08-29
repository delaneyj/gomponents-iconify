package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareWithUpperLeftDiagonalBlack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="m12.5 12l48 1l-48 48V12Z"/><path fill="#fff" d="M60 12v48H12l48-48Z"/><path fill="#3F3F3F" d="m12.5 12l48 1l-48 48V12Z"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-linejoin="round" d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07a.968.968 0 0 1 .965-.965h46.07a.968.968 0 0 1 .965.965v46.07a.968.968 0 0 1-.965.965Z"/><path stroke-linecap="round" d="M59 13L15 57"/></g>`),
		g.Group(children),
	)
}