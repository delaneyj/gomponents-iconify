package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseSymbolForBeginner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#66757F" d="M17 36a1.99 1.99 0 0 1-1.414-.586l-11-11A2 2 0 0 1 4 23V2A2 2 0 0 1 7.414.586L17 10.171L26.586.586A1.998 1.998 0 0 1 30 2v21a2 2 0 0 1-.586 1.414l-11 11A1.99 1.99 0 0 1 17 36z"/><path fill="#47DED4" d="M17 13L28 2v21L17 34z"/><path fill="#FFFF87" d="m6 2l11 11v21L6 23z"/>`),
		g.Group(children),
	)
}