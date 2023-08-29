package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdCheckmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M186.301 339.893L96 249.461l-32 30.507L186.301 402 448 140.506 416 110z" fill="currentColor"/>`),
		g.Group(children),
	)
}