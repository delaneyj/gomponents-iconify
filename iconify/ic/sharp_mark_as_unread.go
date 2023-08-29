package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpMarkAsUnread(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.23 7h4.12L10.5 2L2 6.21V17h2V7.4L10.5 4z"/><path fill="currentColor" d="M5 8v13h17V8H5zm15 4l-6.5 3.33L7 12v-2l6.5 3.33L20 10v2z"/>`),
		g.Group(children),
	)
}