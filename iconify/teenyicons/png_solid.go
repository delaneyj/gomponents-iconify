package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PngSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3 8h.5a.5.5 0 0 0 0-1H3v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM2 6h1.5a1.5 1.5 0 1 1 0 3H3v2H2V6Zm8 0h3v1h-2v3h1V8.5h1V11h-3V6ZM7 8.618V11H6V6h1v.382l1 2V6h1v5H8v-.382l-1-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}