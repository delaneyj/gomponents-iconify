package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M160 115.4L180.7 96 352 256 180.7 416 160 396.7 310.5 256z" fill="currentColor"/>`),
		g.Group(children),
	)
}