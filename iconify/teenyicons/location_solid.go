package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 1.018V0H7v1.018a6.5 6.5 0 0 0-5.981 5.977H0v1h1.019A6.508 6.508 0 0 0 7 13.981V15h1v-1.019a6.508 6.508 0 0 0 5.981-5.986H15v-1h-1.019A6.5 6.5 0 0 0 8 1.018ZM8 3v3.995h4v1H8V12H7V7.995H3v-1h4V3h1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}