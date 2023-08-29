package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosLocation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M256 32c-79.529 0-144 64.471-144 144 0 112 144 304 144 304s144-192 144-304c0-79.529-64.471-144-144-144zm0 190.9c-25.9 0-46.9-21-46.9-46.9s21-46.9 46.9-46.9 46.9 21 46.9 46.9-21 46.9-46.9 46.9z" fill="currentColor"/>`),
		g.Group(children),
	)
}