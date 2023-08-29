package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleBig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 1.5a8.5 8.5 0 1 0 0 17a8.5 8.5 0 0 0 0-17ZM.5 10a9.5 9.5 0 1 1 19 0a9.5 9.5 0 0 1-19 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}