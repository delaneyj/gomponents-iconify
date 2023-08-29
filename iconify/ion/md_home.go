package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M208 448V320h96v128h97.6V256H464L256 64 48 256h62.4v192z" fill="currentColor"/>`),
		g.Group(children),
	)
}