package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkLockedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.875 22q-.35 0-.6-.25t-.25-.6v-3.3q0-.35.25-.6t.6-.25h.15v-1q0-.825.588-1.413T20.024 14q.825 0 1.413.588T22.025 16v1h.15q.35 0 .6.25t.25.6v3.3q0 .35-.25.6t-.6.25h-4.3Zm1.15-5h2v-1q0-.425-.288-.713T20.025 15q-.425 0-.712.288t-.288.712v1ZM4.45 22q-.675 0-.938-.613t.213-1.087l16.6-16.6q.475-.475 1.088-.213t.612.938V12h-2q-2.075 0-3.537 1.463T15.024 17v5H4.45Z"/>`),
		g.Group(children),
	)
}