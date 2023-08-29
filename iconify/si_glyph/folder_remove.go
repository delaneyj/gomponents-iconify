package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.35 5L6.788 2.042H3.021v1.021H1.023v9.913h1.02l.002 1h14.902L16.968 5H8.35zm2.733 5.042H6.979V8.917h4.104v1.125z"/><path d="M14.964 3.982v-.94h-5.94l.33.94h5.61z"/></g>`),
		g.Group(children),
	)
}