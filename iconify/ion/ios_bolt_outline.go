package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosBoltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M273.9 132l-16.5 89.1L254 240h84l-99.8 140 16.5-89.1L258 272h-83.9l99.8-140m28.8-68L143 288h95.8l-29.5 160L369 224h-95.8l29.5-160z" fill="currentColor"/>`),
		g.Group(children),
	)
}