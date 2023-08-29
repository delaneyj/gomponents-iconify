package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserData(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 8h2v8h-2zm-5-3h2v11h-2zm-5 5h2v6h-2zm-2 20h-2v-6a3.003 3.003 0 0 0-3-3H7a3.003 3.003 0 0 0-3 3v6H2v-6a5.006 5.006 0 0 1 5-5h4a5.006 5.006 0 0 1 5 5zM9 9a3 3 0 1 1-3 3a3 3 0 0 1 3-3m0-2a5 5 0 1 0 5 5a5 5 0 0 0-5-5z"/>`),
		g.Group(children),
	)
}