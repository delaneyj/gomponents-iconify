package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadingHOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 7h2v10h-2v-4H7v4H5V7h2v4h4V7zm6.57 0A4.742 4.742 0 0 1 15 9v1h2v7h2V7h-1.43z"/>`),
		g.Group(children),
	)
}