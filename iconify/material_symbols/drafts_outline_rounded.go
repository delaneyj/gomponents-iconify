package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DraftsOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.025 1.6l8.025 4.8q.45.275.7.75t.25 1V19q0 .825-.588 1.413T20 21H4q-.825 0-1.413-.588T2 19V8.15q0-.525.25-1t.7-.75l8.025-4.8q.475-.275 1.025-.275t1.025.275ZM12 12.65L19.8 8L12 3.35L4.2 8l7.8 4.65Zm-1.025 1.725L4 10.2V19h16v-8.8l-6.975 4.175q-.475.275-1.025.275t-1.025-.275ZM13.025 19H20H4h9.025Z"/>`),
		g.Group(children),
	)
}