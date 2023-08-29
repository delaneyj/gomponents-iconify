package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentNext(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.5A1.5 1.5 0 0 0 1.5 4v6A1.5 1.5 0 0 0 3 11.5h1.5v1.6l2.067-1.462l.194-.138H13a1.5 1.5 0 0 0 1.5-1.5V4A1.5 1.5 0 0 0 13 2.5H3ZM0 4a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H7.239l-3.056 2.162L3 16v-3a3 3 0 0 1-3-3V4Zm9.191 3.75H4.75a.75.75 0 0 1 0-1.5h4.441a.75.75 0 0 1 1.09-1.03l1.249 1.25l.53.53l-.53.53l-1.25 1.25a.75.75 0 0 1-1.089-1.03Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}