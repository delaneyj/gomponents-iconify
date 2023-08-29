package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M3 277q0 28 18 46t46 18h170q28 0 46-18t18-46V107q0-28-18-46t-46-18H67q-28 0-46 18T3 107v170zm42-170q0-22 22-22h170q22 0 22 22v170q0 22-22 22H67q-22 0-22-22V107zm171 42H88q0-1 16 20.5t32 43.5l16 22z"/>`),
		g.Group(children),
	)
}