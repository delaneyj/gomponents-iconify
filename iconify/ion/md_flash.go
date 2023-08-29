package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdFlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M160 48v224h64v192l128-256h-64l64-160H160z" fill="currentColor"/>`),
		g.Group(children),
	)
}