package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.35 5L6.788 2.042H3.021v1.021H1.023v9.913h1.02l.002 1h14.902L16.968 5H8.35zm3.568 6.109l-.828.829l-1.578-1.577l-1.578 1.577l-.83-.829l1.578-1.578l-1.578-1.578l.83-.828l1.578 1.576l1.576-1.576l.83.828l-1.578 1.578l1.578 1.578z"/><path d="M14.964 3.982v-.94h-5.94l.33.94h5.61z"/></g>`),
		g.Group(children),
	)
}