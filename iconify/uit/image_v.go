package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 2h-15A2.502 2.502 0 0 0 2 4.5v15A2.502 2.502 0 0 0 4.5 22h15a2.502 2.502 0 0 0 2.5-2.5v-15A2.502 2.502 0 0 0 19.5 2zm-15 19A1.5 1.5 0 0 1 3 19.5v-5.225l3.763-3.762a1.753 1.753 0 0 1 2.474 0l10.468 10.466c-.068.01-.135.02-.205.021h-15zM21 19.5c0 .378-.145.72-.377.984l-6.916-6.916l1.056-1.055a1.753 1.753 0 0 1 2.474 0L21 16.275V19.5zm0-4.639l-3.056-3.055a2.753 2.753 0 0 0-3.888 0L13 12.86L9.944 9.806a2.815 2.815 0 0 0-3.888 0L3 12.86V4.5A1.5 1.5 0 0 1 4.5 3h15A1.5 1.5 0 0 1 21 4.5v10.361z"/>`),
		g.Group(children),
	)
}