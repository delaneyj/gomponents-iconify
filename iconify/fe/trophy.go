package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trophy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feTrophy0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTrophy1" fill="currentColor"><path id="feTrophy2" d="M6.207 9H6a2 2 0 1 1 0-4V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v1a2 2 0 1 1 0 4h-.207A5.504 5.504 0 0 1 13 12.978V17h1a2 2 0 0 1 2 2v1h1a1 1 0 0 1 0 2H7a1 1 0 0 1 0-2h1v-1a2 2 0 0 1 2-2h1v-4.022A5.504 5.504 0 0 1 6.207 9ZM8 4v3.5a3.5 3.5 0 0 0 3.5 3.5h1A3.5 3.5 0 0 0 16 7.5V4H8Z"/></g></g>`),
		g.Group(children),
	)
}