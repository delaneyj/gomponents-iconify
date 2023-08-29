package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 13c1.1 0 2-.9 2-2s-.9-2-2-2h-.1c0-.3.1-.6.1-1c0-2.2-1.8-4-4-4c-.8 0-1.5.2-2.2.6C7.5 3.7 6.6 3 5.5 3C4.1 3 3 4.1 3 5.5c0 .6.2 1.1.6 1.6C3.4 7 3.2 7 3 7c-1.7 0-3 1.3-3 3s1.3 3 3 3h11z"/>`),
		g.Group(children),
	)
}