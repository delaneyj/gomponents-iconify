package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XlsSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12Zm2 5.793V6H2v1.707l.793.793L2 9.293V11h1V9.707l.5-.5l.5.5V11h1V9.293L4.207 8.5L5 7.707V6H4v1.293l-.5.5l-.5-.5ZM6 6h1v4h2v1H6V6Zm7 0h-3v3h2v1h-2v1h3V8h-2V7h2V6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}