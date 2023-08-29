package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiAddOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 21v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2Zm-6 0L0 9q2.375-2.425 5.488-3.713T12 4q3.4 0 6.513 1.288T24 9l-2.425 2.425q-.525-.25-1.1-.388t-1.2-.162L21.1 9.05q-1.975-1.5-4.3-2.275T12 6q-2.475 0-4.8.775T2.9 9.05l9.1 9.1l.875-.875q.025.625.162 1.2t.388 1.1L12 21Zm0-8.925Z"/>`),
		g.Group(children),
	)
}