package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bloodtype(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-3.425 0-5.713-2.35T4 13.8q0-2.5 1.988-5.438T12 2q4.025 3.425 6.013 6.363T20 13.8q0 3.5-2.288 5.85T12 22Zm-3-4h6v-2H9v2Zm2-3h2v-2h2v-2h-2V9h-2v2H9v2h2v2Z"/>`),
		g.Group(children),
	)
}