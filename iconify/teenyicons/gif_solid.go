package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GifSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM2 11V6h3v1H3v3h1V8.5h1V11H2Zm6-4h1V6H6v1h1v3H6v1h3v-1H8V7Zm5-1v1h-2v1h1v1h-1v2h-1V6h3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}