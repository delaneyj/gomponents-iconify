package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareWithLowerRightDiagonalBlack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M59 11.5L60 59H11.5L59 11.5Z"/><path fill="#fff" d="M12 12h48L12 60V12Z"/><path fill="#3F3F3F" d="M59 11.5L60 59H11.5L59 11.5Z"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-linejoin="round" d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07a.968.968 0 0 1 .965-.965h46.07a.968.968 0 0 1 .965.965v46.07a.968.968 0 0 1-.965.965Z"/><path stroke-linecap="round" d="M59 13L15 57"/></g>`),
		g.Group(children),
	)
}