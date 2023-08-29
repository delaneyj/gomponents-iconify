package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SentimentContentOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 15.5h6V14H9v1.5Zm-.5-7q-.8 0-1.488.45T5.876 10.1l1.25.825q.25-.375.6-.638t.775-.262q.425 0 .775.263t.6.612l1.25-.825q-.45-.675-1.137-1.125T8.5 8.5Zm7 0q-.8 0-1.488.45t-1.137 1.15l1.25.825q.25-.35.6-.613t.775-.262q.425 0 .788.25t.587.625l1.25-.825q-.45-.7-1.137-1.15T15.5 8.5ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Zm0-10Zm0 8q3.35 0 5.675-2.325T20 12q0-3.35-2.325-5.675T12 4Q8.65 4 6.325 6.325T4 12q0 3.35 2.325 5.675T12 20Z"/>`),
		g.Group(children),
	)
}