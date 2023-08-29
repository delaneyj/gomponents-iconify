package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderWarning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.35 5L6.788 2.042H3.021v1.021H1.023v9.913h1.021l.001 1h14.902L16.968 5H8.35zm-3.329 8.021l4.5-7.041l4.459 7.041H5.021z"/><path d="M14.964 3.982v-.94h-5.94l.33.94h5.61zM9 8h1v1.956H9zm0 3h1v.973H9z"/></g>`),
		g.Group(children),
	)
}