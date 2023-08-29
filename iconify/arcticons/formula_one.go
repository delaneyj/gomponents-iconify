package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormulaOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.5 18.306h-5.922L27.189 29.694h5.922L44.5 18.306z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.578 18.306h-20.79a7 7 0 0 0-4.95 2.05L3.5 29.694h5.922l6.568-6.567a3.5 3.5 0 0 1 2.475-1.025h16.317"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.422 29.694h5.922l3.357-3.356a1.5 1.5 0 0 1 1.061-.44h11.223"/>`),
		g.Group(children),
	)
}