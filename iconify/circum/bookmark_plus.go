package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.6 21.938a1.482 1.482 0 0 1-1.011-.4l-4.251-3.9a.5.5 0 0 0-.678 0l-4.25 3.9a1.5 1.5 0 0 1-2.517-1.1V4.563a2.5 2.5 0 0 1 2.5-2.5h9.214a2.5 2.5 0 0 1 2.5 2.5v15.872a1.483 1.483 0 0 1-.9 1.375a1.526 1.526 0 0 1-.607.128ZM12 16.5a1.5 1.5 0 0 1 1.018.395l4.251 3.905a.5.5 0 0 0 .838-.368V4.563a1.5 1.5 0 0 0-1.5-1.5H7.393a1.5 1.5 0 0 0-1.5 1.5v15.872a.5.5 0 0 0 .839.368l4.251-3.903A1.5 1.5 0 0 1 12 16.5Z"/><path fill="currentColor" d="M14 10.28h-1.5v1.5a.5.5 0 0 1-1 0v-1.5H10a.5.5 0 0 1 0-1h1.5v-1.5a.5.5 0 0 1 1 0v1.5H14a.5.5 0 0 1 0 1Z"/>`),
		g.Group(children),
	)
}