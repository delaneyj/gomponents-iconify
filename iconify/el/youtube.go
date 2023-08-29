package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Youtube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M1200 1055.438H0V144.562h1200v910.876zm-772.708-189.34l419.616-263.539l-419.616-263.54v527.079z"/>`),
		g.Group(children),
	)
}