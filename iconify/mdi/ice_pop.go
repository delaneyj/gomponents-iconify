package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IcePop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.15 14.96l-8.2 3.69l-4.1-9a3.603 3.603 0 0 0 2.3-3.29c-.01-1.36-.79-2.6-2-3.21c.39-.35.85-.65 1.3-.9c2.26-1 4.92-.02 6 2.21m-.3 13.9l1.6 3.5l2.7-1.21l-1.6-3.5"/>`),
		g.Group(children),
	)
}