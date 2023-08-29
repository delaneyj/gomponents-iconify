package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Videocamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M4.5 34.5h24c2.529 0 3-.471 3-3v-9l7.52 10.18c1.15 1.12 1.91 1.15 2.48-.25V10.61c-.57-1.4-1.33-1.44-2.48-.32L31.5 21.5v-9c0-2.5-.48-3-3-3h-24c-2.48 0-3 .55-3 3.45V31.5c0 2.5.48 3 3 3z"/>`),
		g.Group(children),
	)
}