package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 6H4a2.002 2.002 0 0 0-2 2v16a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2V8a2.002 2.002 0 0 0-2-2Zm-8 18h-8v-2a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1Zm8 0h-6v-2a3 3 0 0 0-3-3h-6a3 3 0 0 0-3 3v2H4V8h24Z"/><path fill="currentColor" d="M16 18a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm0-6a2 2 0 1 1-2 2a2 2 0 0 1 2-2Z"/>`),
		g.Group(children),
	)
}