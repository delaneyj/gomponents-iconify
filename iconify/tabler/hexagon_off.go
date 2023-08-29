package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8.693 4.69l2.336-1.39a2.056 2.056 0 0 1 2 0l6 3.573H19a2 2 0 0 1 1 1.747v6.536c0 .246-.045.485-.13.707m-2.16 1.847l-4.739 3.027a2 2 0 0 1-1.942 0l-6-3.833A2 2 0 0 1 4 15.157V8.62a2 2 0 0 1 1.029-1.748l1.154-.687M3 3l18 18"/>`),
		g.Group(children),
	)
}