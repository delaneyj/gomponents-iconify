package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DryCleaningRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 22q-.825 0-1.413-.588T7 20v-4H5.4q-.975 0-1.688-.713T3 13.6q0-.725.388-1.337t1.062-.913L11 8.45V7.8q-.875-.3-1.438-1.062T9 5q0-1.275.863-2.138T12 2q1.05 0 1.813.513t.987 1.312q.125.475-.15.825t-.825.35q-.35 0-.6-.163t-.45-.437q-.125-.2-.338-.3T12 4q-.425 0-.713.288T11 5q0 .425.288.713T12 6t.713.288Q13 6.574 13 7v1.45l6.55 2.9q.675.3 1.063.913T21 13.6q0 .975-.713 1.688T18.6 16H17v4q0 .825-.588 1.413T15 22H9Zm-3.6-8h1.875q.325-.5.738-.75T9 13h6q.575 0 .988.25t.737.75H18.6q.225 0 .313-.15t.087-.3q0-.125-.063-.213t-.187-.137l-6.75-3l-6.75 3q-.125.05-.188.138T5 13.55q0 .2.1.325t.3.125Z"/>`),
		g.Group(children),
	)
}