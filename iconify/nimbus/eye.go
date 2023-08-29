package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 5.5A2.59 2.59 0 0 0 5.33 8A2.59 2.59 0 0 0 8 10.5A2.59 2.59 0 0 0 10.67 8A2.59 2.59 0 0 0 8 5.5zm0 3.75A1.35 1.35 0 0 1 6.58 8A1.35 1.35 0 0 1 8 6.75A1.35 1.35 0 0 1 9.42 8A1.35 1.35 0 0 1 8 9.25z"/><path fill="currentColor" d="M8 2.5A8.11 8.11 0 0 0 0 8a8.11 8.11 0 0 0 8 5.5A8.11 8.11 0 0 0 16 8a8.11 8.11 0 0 0-8-5.5zm5.4 7.5A6.91 6.91 0 0 1 8 12.25A6.91 6.91 0 0 1 2.6 10a7.2 7.2 0 0 1-1.27-2A7.2 7.2 0 0 1 2.6 6A6.91 6.91 0 0 1 8 3.75A6.91 6.91 0 0 1 13.4 6a7.2 7.2 0 0 1 1.27 2a7.2 7.2 0 0 1-1.27 2z"/>`),
		g.Group(children),
	)
}