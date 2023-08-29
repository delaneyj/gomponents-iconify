package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2h6v2h-6v6h-2V4h2V2zm0 8h6v2h-6v-2zm8-6h-2v6h2V4zM7 16v6h16v-6h-2v4H9v-4h12v-2H9v2H7zm-1-6H4V8H2v2h2v2H2v2h2v-2h2v2h2v-2H6v-2zm0 0h2V8H6v2z"/>`),
		g.Group(children),
	)
}