package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yacht(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 19.216L5.414 24h15.882l2.274-3.195L.002 19.216zM13.985 0v19.073l7.309.503zm-.996 3.214L2.429 18.241l10.56.713z"/>`),
		g.Group(children),
	)
}