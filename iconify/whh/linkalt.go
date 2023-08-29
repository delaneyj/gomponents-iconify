package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Linkalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M947.533 947.5q-76.5 76.5-185 76.5t-184.5-77l-51-51l113-64l24 24q40 40 96 40t96-40t40-96t-40-96l-240-241q-40-39-96.5-39t-96.5 39q-1 2-3 4q-31 31-36 53l-96-96q4-7 10-15t9.5-11.5l13-13l12.5-11.5q76-77 184.5-77t184.5 77l245 245q77 76 77 184.5t-76.5 185zM360.033 168q-40-40-96-40t-96 40t-40 96t40 96l240 241q40 40 96.5 40t96.5-40q1-2 3-4q31-31 36-53l96 96q-4 7-10 15t-9.5 11.5l-13 13l-12.5 11.5q-76 77-184.5 77t-184.5-77l-245-245q-77-76-77-184.5t76.5-185t185-76.5t184.5 77l51 51l-113 64z"/>`),
		g.Group(children),
	)
}