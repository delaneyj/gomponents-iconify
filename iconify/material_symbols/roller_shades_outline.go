package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RollerShadesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21v-2h2V3h16v16h2v2H2Zm4-10h12V5H6v6Zm0 8h12v-6h-5v1.8q.35.25.55.625t.2.825q0 .725-.513 1.238T12 18q-.725 0-1.238-.513t-.512-1.237q0-.45.2-.813t.55-.612V13H6v6ZM6 5h12H6Z"/>`),
		g.Group(children),
	)
}