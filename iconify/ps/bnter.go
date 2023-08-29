package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bnter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M38 321h17q4 21-9 43q5 0 13.5-2t25-12.5T105 321h321q15 0 25.5-10.5T462 285V57q0-15-10.5-26T426 20H38q-15 0-25.5 11T2 57v228q0 15 10.5 25.5T38 321zm350-125q7 0 12 5t5 12t-5 12t-12 5t-12-5t-5-12q0-17 17-17zm-104 52q1 1 1 4q0 18-13 30.5T241 295q-14 0-25.5-8.5T200 265zM78 196q17 0 17 17q0 7-5 12t-12 5t-12-5t-5-12t5-12t12-5z"/>`),
		g.Group(children),
	)
}