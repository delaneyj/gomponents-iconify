package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SvgSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM5 6H2v3h2v1H2v1h3V8H3V7h2V6Zm2 0H6v3.707l1.5 1.5l1.5-1.5V6H8v3.293l-.5.5l-.5-.5V6Zm3 0h3v1h-2v3h1V8.5h1V11h-3V6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}