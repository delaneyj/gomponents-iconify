package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleBig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 2.5a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15ZM.5 10a9.5 9.5 0 1 1 19 0a9.5 9.5 0 0 1-19 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}