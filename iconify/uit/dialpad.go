package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dialpad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 10H3a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4A.5.5 0 0 0 7 10zm-.5 4h-3v-3h3v3zM7 3H3a.5.5 0 0 0-.5.5v4A.5.5 0 0 0 3 8h4a.5.5 0 0 0 .5-.5v-4A.5.5 0 0 0 7 3zm-.5 4h-3V4h3v3zM14 3h-4a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4A.5.5 0 0 0 14 3zm-.5 4h-3V4h3v3zM21 3h-4a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4A.5.5 0 0 0 21 3zm-.5 4h-3V4h3v3zM14 17h-4a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5zm-.5 4h-3v-3h3v3zM21 10h-4a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5zm-.5 4h-3v-3h3v3zM14 10h-4a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5zm-.5 4h-3v-3h3v3z"/>`),
		g.Group(children),
	)
}