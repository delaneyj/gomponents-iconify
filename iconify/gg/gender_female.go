package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderFemale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 3a5 5 0 0 0-1 9.9V15H8v2h3v4h2v-4h3v-2h-3v-2.1A5.002 5.002 0 0 0 12 3ZM9 8a3 3 0 1 0 6 0a3 3 0 0 0-6 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}