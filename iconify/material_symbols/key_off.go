package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.75 22.575l-7.55-7.55q-.8 1.35-2.175 2.163T7 18q-2.5 0-4.25-1.75T1 12q0-1.65.813-3.025T3.974 6.8l-2.55-2.55l1.425-1.4l18.3 18.325l-1.4 1.4ZM21.025 10L23 11.975l-4 4L17 14l-.075.1l-2.1-2.1l-2-2h8.2ZM7 15q1.075 0 1.875-.663T9.9 12.726l-.563-.563l-1.25-1.25l-1.25-1.25l-.562-.562q-1.05.225-1.663 1.063T4 12q0 1.25.875 2.125T7 15Z"/>`),
		g.Group(children),
	)
}