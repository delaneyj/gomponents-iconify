package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m12 8.954l5.635-1.127a2.942 2.942 0 0 0-1.154-5.769l-4.07.814A3 3 0 0 0 10 5.814v8.076a4 4 0 1 0 2 3.465v-8.4Zm4.874-4.935l-4.07.814a1 1 0 0 0-.804.98v1.102l5.243-1.049a.942.942 0 0 0-.37-1.847ZM10 17.354a2 2 0 1 0-4 0a2 2 0 0 0 4 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}