package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func F(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M255 5v78c-11-7-25-11-39-11c-39 0-70 31-71 70v86h110v60H145v466H73V288H0v-60h73v-84c0-53 29-99 72-124c22-12 45-20 71-20c14 0 27 2 39 5z"/>`),
		g.Group(children),
	)
}