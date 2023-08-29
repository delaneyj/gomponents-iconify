package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosBolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M302.7 64L143 288h95.8l-29.5 160L369 224h-95.8l29.5-160z" fill="currentColor"/>`),
		g.Group(children),
	)
}