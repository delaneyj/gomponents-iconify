package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drafts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1l9.05 5.4q.45.275.7.75t.25 1V19q0 .825-.588 1.413T20 21H4q-.825 0-1.413-.588T2 19V8.15q0-.525.25-1t.7-.75L12 1Zm0 11.65L19.8 8L12 3.35L4.2 8l7.8 4.65Z"/>`),
		g.Group(children),
	)
}