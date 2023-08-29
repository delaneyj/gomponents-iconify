package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoStoriesOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20q-1.2-.95-2.6-1.475T6.5 18q-1.475 0-2.85.463T1 19.674V5.55q1.2-.85 2.613-1.2T6.5 4q1.45 0 2.85.375T12 5.5v12.1q1.25-.775 2.638-1.188T17.5 16q.9 0 1.775.15T21 16.6v-12q.525.175 1.038.4t.962.55v14.125q-1.275-.75-2.65-1.213T17.5 18q-1.5 0-2.9.525T12 20Zm2-5V5.5l5-5v10L14 15Zm-4 1.625v-9.9q-.825-.35-1.713-.537T6.5 6q-.9 0-1.788.175T3 6.7v9.925q.85-.325 1.725-.475T6.5 16q.9 0 1.775.15t1.725.475Zm0 0v-9.9v9.9Z"/>`),
		g.Group(children),
	)
}