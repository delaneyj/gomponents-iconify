package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M2.68 15.726c-.1.039-.16.09-.18.149l.18-.149zm-.18.149v20.627c0 2.509.49 2.998 3 2.998h7c2.51 0 3-.461 3-3v-8h9v8.031c0 2.51.51 2.979 3 2.969c.04.031 7 0 7 0c2.529 0 3-.526 3-2.997V16.495c0-.08-.09-.15-.27-.23L20 1.5L2.68 15.726l-.18.149z"/>`),
		g.Group(children),
	)
}