package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosPlayOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M144 124.9L353.8 256 144 387.1V124.9M128 96v320l256-160L128 96z" fill="currentColor"/>`),
		g.Group(children),
	)
}