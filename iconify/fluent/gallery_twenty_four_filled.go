package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GalleryTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M5.25 3A2.25 2.25 0 0 0 3 5.25v6h8.25V3h-6z" fill="currentColor"/><path d="M12.75 3v8.25H21v-6A2.25 2.25 0 0 0 18.75 3h-6z" fill="currentColor"/><path d="M21 12.75h-8.25V21h6A2.25 2.25 0 0 0 21 18.75v-6z" fill="currentColor"/><path d="M11.25 21v-8.25H3v6A2.25 2.25 0 0 0 5.25 21h6z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}