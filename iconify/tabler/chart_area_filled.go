package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartAreaFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M20 18a1 1 0 0 1 .117 1.993L20 20H4a1 1 0 0 1-.117-1.993L4 18h16zM15.22 5.375a1 1 0 0 1 1.393-.165l.094.083l4 4a1 1 0 0 1 .284.576L21 10v5a1 1 0 0 1-.883.993L20 16H3.978l-.11-.009l-.11-.02l-.107-.034l-.105-.046l-.1-.059l-.094-.07l-.06-.055l-.072-.082l-.064-.089l-.054-.096l-.016-.035l-.04-.103l-.027-.106l-.015-.108l-.004-.11l.009-.11l.019-.105c.01-.04.022-.077.035-.112l.046-.105l.059-.1l4-6a1 1 0 0 1 1.165-.39l.114.05l3.277 1.638l3.495-4.369z"/></g>`),
		g.Group(children),
	)
}