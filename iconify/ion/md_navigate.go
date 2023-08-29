package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdNavigate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M256 64L96 433.062 110.938 448 256 384l145.062 64L416 433.062z" fill="currentColor"/>`),
		g.Group(children),
	)
}