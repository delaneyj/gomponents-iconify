package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandCoinFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.335 11.502h2.17a4.5 4.5 0 0 1 4.5 4.5H9.004v1h8v-1a5.578 5.578 0 0 0-.885-3h2.886a5 5 0 0 1 4.516 2.852c-2.365 3.121-6.194 5.149-10.516 5.149c-2.761 0-5.1-.59-7-1.625v-9.304a6.966 6.966 0 0 1 3.33 1.428Zm-4.33 7.5a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v9Zm13-14a3 3 0 1 1 0 6.001a3 3 0 0 1 0-6Zm-7-3a3 3 0 1 1 0 6.001a3 3 0 0 1 0-6Z"/>`),
		g.Group(children),
	)
}