package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25.98 12.896h-6.67V6.23h-6.665v6.666H5.98v6.666h6.667v6.667h6.665V19.56h6.667z"/>`),
		g.Group(children),
	)
}