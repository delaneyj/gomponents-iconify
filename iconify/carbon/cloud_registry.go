package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudRegistry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 30h-5a2.002 2.002 0 0 1-2-2v-5a2.002 2.002 0 0 1 2-2h5a2.002 2.002 0 0 1 2 2v5a2.002 2.002 0 0 1-2 2zm-5-7h-.001L23 28h5v-5zm-7 0h-5a2.002 2.002 0 0 1-2-2V11a2.002 2.002 0 0 1 2-2h10a2.002 2.002 0 0 1 2 2v5h-2v-5H11v10h5z"/><path fill="currentColor" d="M16 30H4a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v12h-2V4H4v24h12Z"/>`),
		g.Group(children),
	)
}