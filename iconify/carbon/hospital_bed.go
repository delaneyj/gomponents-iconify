package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalBed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 16h-8a2.002 2.002 0 0 0-2 2v6H4V14H2v16h2v-4h24v4h2v-9a5.006 5.006 0 0 0-5-5Zm3 8H17v-6h8a3.003 3.003 0 0 1 3 3Z"/><path fill="currentColor" d="M9.5 17A1.5 1.5 0 1 1 8 18.5A1.502 1.502 0 0 1 9.5 17m0-2a3.5 3.5 0 1 0 3.5 3.5A3.5 3.5 0 0 0 9.5 15zM21 6h-4V2h-2v4h-4v2h4v4h2V8h4V6z"/>`),
		g.Group(children),
	)
}