package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FindReplace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.1 10q.35-2.575 2.288-4.288T11 4q1.475 0 2.763.563T16 6.1V4h2v6h-6V8h3q-.725-.9-1.738-1.45T11 6Q9.2 6 7.825 7.137T6.1 10h-2Zm15.5 11l-4.4-4.4q-.9.675-1.962 1.038T11 18q-1.475 0-2.763-.563T6 15.9V18H4v-6h6v2H7q.725.9 1.738 1.45T11 16q1.8 0 3.175-1.137T15.9 12h2q-.125.9-.45 1.688T16.6 15.2l4.4 4.4l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}