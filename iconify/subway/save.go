package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Save(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0C128 149.3 0 42.7 0 42.7V128c0 170.7 149.3 384 256 384s256-213.3 256-384V42.7S384 149.3 256 0z"/>`),
		g.Group(children),
	)
}