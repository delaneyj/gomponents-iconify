package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brightness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M4 4h4.686L12 .686L15.314 4H20v4.686L23.314 12L20 15.314V20h-4.686L12 23.314L8.686 20H4v-4.686L.686 12L4 8.686V4zm8 2v12a6 6 0 0 0 0-12z" fill="currentColor"/>`),
		g.Group(children),
	)
}