package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LungsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 5.5c1.412.47 2.048 2.159 2.327 4.023l-4.523 2.611l1 1.732l3.71-2.141C11.06 13.079 11 14.309 11 15c0 3-1 6-5 6s-4 0-4-4C2 9.5 5.5 4.5 8.5 5.5ZM22.001 17v.436c-.005 3.564-.15 3.564-4 3.564c-4 0-5-3-5-6c0-.691-.06-1.92-.014-3.274l3.71 2.14l1-1.732l-4.523-2.61c.279-1.865.915-3.553 2.327-4.024c3-1 6.5 4 6.5 11.5ZM13 2v9h-2V2h2Z"/>`),
		g.Group(children),
	)
}