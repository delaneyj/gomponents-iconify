package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderRss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.35 5L6.788 2.042H3.021v1.021H1.023v9.913h1.021l.001 1h14.902L16.968 5H8.35zm-.327 7.042h-1v-1h.17a.83.83 0 0 1 .83.83v.17zm.97-.06c0-.94-1.018-2.023-1.966-2.023v-.95c1.54 0 2.948 1.445 2.948 2.973h-.982zm2.137.013c0-2.476-1.654-4.115-4.12-4.115v-.838c3.299 0 4.932 1.636 4.932 4.954h-.812z"/><path d="M14.964 3.982v-.94h-5.94l.33.94h5.61z"/></g>`),
		g.Group(children),
	)
}