package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DonutSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.025 2.05q3.575.35 6.088 2.863T21.974 11h-7.15q-.225-.65-.687-1.137t-1.113-.713v-7.1Zm2 2.55V8q.275.225.525.475t.475.525h3.4q-.6-1.5-1.75-2.65t-2.65-1.75Zm-4-2.55v7.1q-.9.325-1.45 1.113T9.025 12q0 .95.55 1.713t1.45 1.087v7.15q-3.85-.375-6.425-3.225T2.025 12q0-3.875 2.575-6.725t6.425-3.225Zm-2 2.55q-2.275.875-3.638 2.9T4.026 12q0 2.475 1.363 4.5t3.637 2.95V16q-.95-.725-1.475-1.763T7.025 12q0-1.2.525-2.238T9.025 8V4.6Zm5.8 8.4h7.15q-.35 3.575-2.863 6.088t-6.087 2.862V14.8q.65-.225 1.113-.688T14.825 13Zm1.2 2q-.2.275-.462.525t-.538.475v3.4q1.5-.6 2.65-1.75t1.75-2.65h-3.4Zm-9-2.975Zm9-3.025Zm0 6Z"/>`),
		g.Group(children),
	)
}