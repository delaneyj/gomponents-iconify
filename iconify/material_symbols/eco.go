package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eco(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.25 22q-.825 0-1.663-.188t-1.712-.537q.3-3.025 1.75-5.65T13.35 11q-2.75 1.4-4.763 3.7t-2.812 5.25q-.1-.075-.188-.162L5.4 19.6q-1.175-1.175-1.788-2.625T3 13.95q0-1.7.675-3.25T5.55 7.95q2.025-2.025 5.25-2.637t9.05-.113q.45 5.975-.15 9.113t-2.6 5.137q-1.225 1.225-2.738 1.888T11.25 22Z"/>`),
		g.Group(children),
	)
}