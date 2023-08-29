package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m8.05 5.82l5.56 5.66l.89-.87L8.91 4.9a1.18 1.18 0 0 0-.86-.39a1.13 1.13 0 0 0-.85.39l-5.7 5.7l.88.89z"/>`),
		g.Group(children),
	)
}