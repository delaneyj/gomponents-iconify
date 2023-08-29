package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DomesticWorkerAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M16.5 14a4 4 0 1 1 0-8a4 4 0 0 1 0 8Zm0-2a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z" clip-rule="evenodd"/><path d="M10 22a2 2 0 0 1 2-2v4a2 2 0 0 1-2-2Z"/><path fill-rule="evenodd" d="M34.205 27.76a4 4 0 0 1 4.797 2.997l2.999 8.545l-11.692 2.7l-1.05-8.995a4 4 0 0 1 2.997-4.797l-1.747-7.566A2.997 2.997 0 0 1 28 22h-5v17a3 3 0 1 1-6 0v-8h-1v8c0 .701-.24 1.346-.644 1.857A3 3 0 0 1 10 39V27.718c-1.563-.72-4-2.808-4-6.147C6 17.364 9.871 16 10.985 16H28c.56 0 1.086.154 1.535.422l-1.103-4.777a1 1 0 0 1 1.949-.45l3.824 16.565Zm-1.5 2.398l1.95-.45a2 2 0 0 1 2.398 1.5l-5.846 1.349a2 2 0 0 1 1.499-2.399Zm5.003 2.95l-6.256 1.445l.584 4.997l.746-.172l-.675-2.924l1.949-.45l.675 2.924l4.643-1.072l-1.666-4.747ZM12 26.439l-1.163-.537C9.693 25.374 8 23.854 8 21.571c0-1.394.605-2.238 1.308-2.789a4.28 4.28 0 0 1 1.126-.63c.304-.112.494-.142.543-.15c.014-.002.016-.002.008-.002H14v9h6v-9h8a1 1 0 1 1 0 2h-7v19a1 1 0 1 1-2 0V29h-5v10a1 1 0 1 1-2 0V26.437ZM18 21v-3h-2v3h2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}