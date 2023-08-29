package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderZipF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.83 2H17a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H3a3 3 0 0 1-3-3V3a3 3 0 0 1 3-3h5c1.306 0 2.417.835 2.83 2zM14 6v2h2V6h-2zm-2-2v2h2V4h-2zm0 4v2h2V8h-2zm2 2v2h2v-2h-2zm-2 2v2h2v-2h-2z"/>`),
		g.Group(children),
	)
}