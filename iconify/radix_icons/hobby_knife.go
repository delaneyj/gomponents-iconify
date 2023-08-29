package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HobbyKnife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.354 13.354a.5.5 0 0 1-.707 0l-5.25-5.25A.5.5 0 0 1 6.316 8H5a.5.5 0 0 1-.472-.336l-2.4-6.9A.5.5 0 0 1 2.936.23l5.4 4.9a.5.5 0 0 1 .164.37v.317a.501.501 0 0 1 .104.08l5.25 5.25a.5.5 0 0 1 0 .707l-1.5 1.5ZM8.25 6.957l-.793.793L12 12.293l.793-.793L8.25 6.957ZM3.717 2.288L5.355 7h.938L7.5 5.793V5.72L3.717 2.288Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}