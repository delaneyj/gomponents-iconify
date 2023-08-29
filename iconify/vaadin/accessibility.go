package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Accessibility(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10.4 10h-.5c.1.3.1.7.1 1c0 2.2-1.8 4-4 4s-4-1.8-4-4c0-2.1 1.6-3.8 3.7-4l-.2-1C2.9 6.4 1 8.4 1 11c0 2.8 2.2 5 5 5c2.4 0 4.4-1.7 4.9-3.9l-.5-2.1z"/><path fill="currentColor" d="M13.1 13L12 8H7.9l-.2-1H11V6H7.5l-.6-2.5c.9-.1 1.6-.8 1.6-1.7C8.5.8 7.7 0 6.7 0S5 .8 5 1.8c0 .6.3 1.2.8 1.5L7.1 9h4.1l1.2 5H15v-1h-1.9z"/>`),
		g.Group(children),
	)
}