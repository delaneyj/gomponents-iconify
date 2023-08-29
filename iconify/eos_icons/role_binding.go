package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoleBinding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 21h-2a2 2 0 0 1 0-4h2v-2h-2a4 4 0 0 0 0 8h2Zm8-2a4 4 0 0 1-4 4h-2v-2h2a2 2 0 0 0 0-4h-2v-2h2a4 4 0 0 1 4 4Z"/><path fill="currentColor" d="M14 18h4v2h-4zm-7 1a5.989 5.989 0 0 1 .09-1H3v-1.4c0-2 4-3.1 6-3.1a8.548 8.548 0 0 1 1.35.125A5.954 5.954 0 0 1 13 13h5V4a2.006 2.006 0 0 0-2-2h-4.18a2.988 2.988 0 0 0-5.64 0H2a2.006 2.006 0 0 0-2 2v14a2.006 2.006 0 0 0 2 2h5.09A5.989 5.989 0 0 1 7 19ZM9 2a1 1 0 1 1-1 1a1.003 1.003 0 0 1 1-1Zm0 4a3 3 0 1 1-3 3a2.996 2.996 0 0 1 3-3Z"/>`),
		g.Group(children),
	)
}