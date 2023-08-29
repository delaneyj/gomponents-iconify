package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstAid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7.527 4.5H.5v.25l.055.31a45.684 45.684 0 0 1 0 15.88l-.055.31v.25h23v-.25l-.055-.31a45.686 45.686 0 0 1 0-15.88l.055-.31V4.5h-7.027m-8.946 0a4.5 4.5 0 0 1 8.945 0m-8.945 0h8.945m-5.972 4v3h-3v3h3v3h3v-3h3v-3h-3v-3h-3Z"/>`),
		g.Group(children),
	)
}