package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.243 4.757a6 6 0 0 1 .236 8.236l-8.48 8.492l-8.478-8.492a6 6 0 0 1 6.74-9.553L6.343 7.358l1.414 1.415L12 4.53l-.013-.014l.014.013a5.998 5.998 0 0 1 8.242.228Z"/>`),
		g.Group(children),
	)
}