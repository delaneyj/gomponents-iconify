package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Designbump(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m258 278l-18 16V130h81L182 2L43 130h80v40l-17-12L5 252h51v210h92V252h57l-67-64v-79H82l100-92l100 92h-57v202l-66 59h56v92h92v-92h51z"/>`),
		g.Group(children),
	)
}