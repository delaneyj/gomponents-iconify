package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutosumTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.04 3.804A.5.5 0 0 1 4.5 3.5h11a.5.5 0 1 1 0 1H5.675l4.553 4.732a.5.5 0 0 1 .023.667L5.57 15.5h9.93a.5.5 0 1 1 0 1h-11a.5.5 0 0 1-.384-.82l5.08-6.078L4.14 4.347a.5.5 0 0 1-.1-.543Z"/>`),
		g.Group(children),
	)
}