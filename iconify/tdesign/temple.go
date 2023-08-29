package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Temple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1.586L19.414 9H21v2h-1v9h2v2H2v-2h2v-9H3V9h1.586L12 1.586ZM6 11v9h5v-9H6Zm7 0v9h5v-9h-5Zm3.586-2L12 4.414L7.414 9h9.172Z"/>`),
		g.Group(children),
	)
}