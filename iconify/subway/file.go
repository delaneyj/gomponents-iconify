package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func File(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M320 0v128h128L320 0zm-21.3 0H64v512h384V149.3H298.7V0z"/>`),
		g.Group(children),
	)
}