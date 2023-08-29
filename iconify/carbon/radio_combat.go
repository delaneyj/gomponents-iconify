package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioCombat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 26a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2zm7 0h9v2h-9zm0 4h9v2h-9z"/><path fill="currentColor" d="M20 12h8a2.002 2.002 0 0 0 2-2V6a2.002 2.002 0 0 0-2-2h-8a2.002 2.002 0 0 0-2 2v1h-6a4.005 4.005 0 0 0-4 4v3H6V4H4v10a2.002 2.002 0 0 0-2 2v12a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2V16a2.002 2.002 0 0 0-2-2H10v-3a2.002 2.002 0 0 1 2-2h6v1a2.002 2.002 0 0 0 2 2Zm4-6h4v4h-4Zm-4 0h2v4h-2Zm8 22H4V16h24Z"/>`),
		g.Group(children),
	)
}