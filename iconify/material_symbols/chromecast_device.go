package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChromecastDevice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 14.5v-5h3v5h-3ZM6 17.25q-1.875 0-3.438-1.125T1 13.5v-3Q1 9 2.563 7.875T6 6.75q1.05 0 1.788.238T9.15 7.5q.475.2.913.35T11 8h8v8h-8q-.5 0-.938.15t-.912.35q-.625.275-1.363.513T6 17.25ZM2.5 12.5h1q.2 0 .35-.15T4 12q0-.2-.15-.35t-.35-.15h-1q-.2 0-.35.15T2 12q0 .2.15.35t.35.15Z"/>`),
		g.Group(children),
	)
}