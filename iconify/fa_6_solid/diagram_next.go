package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiagramNext(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 160c0 35.3-28.7 64-64 64H280v64h46.1c21.4 0 32.1 25.9 17 41L273 399c-9.4 9.4-24.6 9.4-33.9 0L169 329c-15.1-15.1-4.4-41 17-41h46v-64H64c-35.3 0-64-28.7-64-64V96c0-35.3 28.7-64 64-64h384c35.3 0 64 28.7 64 64v64zm-64 256v-64h-82.7l.4-.4c18.4-18.4 20.4-43.7 11-63.6H448c35.3 0 64 28.7 64 64v64c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64v-64c0-35.3 28.7-64 64-64h71.3c-9.4 19.9-7.4 45.2 11 63.6l.4.4H64v64h146.7l5.7 5.7c21.9 21.9 57.3 21.9 79.2 0l5.7-5.7H448z"/>`),
		g.Group(children),
	)
}