package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatLineLeftThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 76h-60V32a4 4 0 0 0-6.83-2.83l-96 96a4 4 0 0 0 0 5.66l96 96A4 4 0 0 0 124 224v-44h60a4 4 0 0 0 4-4V80a4 4 0 0 0-4-4Zm-4 96h-60a4 4 0 0 0-4 4v38.34L29.66 128L116 41.66V80a4 4 0 0 0 4 4h60Zm40-92v96a4 4 0 0 1-8 0V80a4 4 0 0 1 8 0Z"/>`),
		g.Group(children),
	)
}