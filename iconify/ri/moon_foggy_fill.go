package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonFoggyFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 20.334V18h-2v-4H3.332A9.512 9.512 0 0 1 3 11.5c0-4.56 3.213-8.37 7.5-9.29a8 8 0 0 0 11.49 9.724a9.505 9.505 0 0 1-5.99 8.4ZM7 20h7v2H7v-2Zm-5-4h10v2H2v-2Z"/>`),
		g.Group(children),
	)
}