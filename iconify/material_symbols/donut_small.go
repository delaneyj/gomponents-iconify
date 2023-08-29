package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DonutSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.825 11q-.225-.65-.687-1.137t-1.113-.713v-7.1q3.575.35 6.088 2.863T21.974 11h-7.15Zm-3.8 10.95q-3.85-.375-6.425-3.225T2.025 12q0-3.875 2.575-6.725t6.425-3.225v7.1q-.9.325-1.45 1.113T9.025 12q0 .95.55 1.713t1.45 1.087v7.15Zm2 0V14.8q.65-.225 1.113-.688T14.825 13h7.15q-.35 3.575-2.863 6.088t-6.087 2.862Z"/>`),
		g.Group(children),
	)
}