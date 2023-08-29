package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowGraphDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M320 384l61.8-61.8-93.5-98.2-107 106.7L32 128l149.3 128 107-112 130.9 140.8L480 224v160z" fill="currentColor"/>`),
		g.Group(children),
	)
}