package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DislikeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.663 20.25a1.4 1.4 0 0 0 1.149-.6l4.232-6.077a5.75 5.75 0 0 0 .997-3.914l-.326-2.961a2.736 2.736 0 0 0-2.396-2.419L14.2 4.027a11.068 11.068 0 0 0-4.8.489a3.845 3.845 0 0 0-2.365 2.24L5.327 11.1a2.818 2.818 0 0 0 3.243 3.78l3.77-.85a.074.074 0 0 1 .041 0c.011.004.023.01.035.023a.083.083 0 0 1 .02.034a.074.074 0 0 1 0 .04l-1.128 4.374a1.4 1.4 0 0 0 1.356 1.749Z"/>`),
		g.Group(children),
	)
}