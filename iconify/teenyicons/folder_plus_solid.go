package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderPlusSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h4.207l2 2H13.5A1.5 1.5 0 0 1 15 4.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5v-10ZM7 11V9H5V8h2V6h1v2h2v1H8v2H7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}