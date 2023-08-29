package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CsvSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM2 6h3v1H3v3h2v1H2V6Zm7 0H6v3h2v1H6v1h3V8H7V7h2V6Zm2 0h-1v3.707l1.5 1.5l1.5-1.5V6h-1v3.293l-.5.5l-.5-.5V6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}