package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M39 14h495c40 0 71 34 71 75v506c0 50-31 85-77 85H43c-31 0-43-24-43-38V51c0-18 14-37 39-37zm95 624h394c22 0 35-16 35-43V89c0-15-12-33-29-33h-47v383l-138-44l-134 44V56h-81v582zm144-422v37h55v56h37v-56h55v-37h-55v-56h-37v56h-55z"/>`),
		g.Group(children),
	)
}