package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessThreeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 12q0 2.025-.738 3.8t-2.037 3.125q-1.3 1.35-3.063 2.163T9.35 22q-.425.025-.713-.288t-.287-.737q0-.25.1-.463t.3-.337q2.05-1.45 3.15-3.575T13 12q0-2.475-1.1-4.6T8.75 3.825q-.2-.125-.3-.338t-.1-.462q0-.425.288-.737T9.35 2q2.05.1 3.813.913t3.062 2.162q1.3 1.35 2.038 3.125T19 12Z"/>`),
		g.Group(children),
	)
}