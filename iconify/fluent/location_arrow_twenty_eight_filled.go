package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationArrowTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M23.959 2.104c1.213-.467 2.405.725 1.938 1.938L17.821 25.04c-.522 1.36-2.48 1.251-2.85-.157l-2.358-8.96a.75.75 0 0 0-.535-.534l-8.96-2.358c-1.408-.37-1.515-2.328-.156-2.85l20.997-8.076Z"/>`),
		g.Group(children),
	)
}