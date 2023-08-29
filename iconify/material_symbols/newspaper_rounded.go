package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewspaperRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19V3.6q0-.175.15-.238t.275.063l.9.9q.15.15.35.15t.35-.15l.95-.975q.15-.15.35-.15t.35.15l.975.975q.15.15.35.15t.35-.15l.975-.975q.15-.15.35-.15t.35.15l.95.975q.15.15.35.15t.35-.15l.975-.975q.15-.15.35-.15t.35.15l.975.975q.15.15.35.15t.35-.15l.95-.975q.15-.15.35-.15t.35.15l.975.975q.15.15.35.15t.35-.15l.975-.975q.15-.15.35-.15t.35.15l.95.975q.15.15.35.15t.35-.15l.9-.9q.125-.125.275-.062T22 3.6V19q0 .825-.588 1.413T20 21H4Zm0-2h7v-6H4v6Zm9 0h7v-2h-7v2Zm0-4h7v-2h-7v2Zm-9-4h16V8H4v3Z"/>`),
		g.Group(children),
	)
}