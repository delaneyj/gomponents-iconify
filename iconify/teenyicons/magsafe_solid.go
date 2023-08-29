package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagsafeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M12 0H3v3h9V0Zm2 4H1v6h3v2h1v3h1v-3h3v3h1v-3h1v-2h3V4Z"/>`),
		g.Group(children),
	)
}