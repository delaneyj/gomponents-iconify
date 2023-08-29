package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tags(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9 2H7.5l7 7l-5.3 5.2l.8.8l6-6z"/><path fill="currentColor" d="M6 2H0v6l7 7l6-6l-7-7zM2.8 6c-.7 0-1.3-.6-1.3-1.2s.6-1.2 1.2-1.2S4 4.1 4 4.8S3.4 6 2.8 6z"/>`),
		g.Group(children),
	)
}