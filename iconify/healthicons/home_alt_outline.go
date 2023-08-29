package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M41.708 23.794L24.058 6.086L6.294 23.792a1 1 0 1 0 1.412 1.416L10 22.922V41a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1v-8.74a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1V41a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V22.907l2.292 2.3a1 1 0 1 0 1.416-1.413ZM36 20.9L24.053 8.914L12 20.928V40h7v-7.74a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3V40h7V20.9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}