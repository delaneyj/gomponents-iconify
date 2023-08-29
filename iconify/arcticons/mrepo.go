package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mrepo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.904 5.502c2.28-.05 4.202.96 6.203 2.121l8.415 4.783c2.57 1.5 3.201 3.401 3.201 6.083v9.545c.04 2.772.57 4.253-2.301 6.564c-2.872 2.311-9.285 6.203-12.717 7.364s-5.383.28-7.754-1.09c-3.242-1.872-6.804-3.603-9.556-5.424s-3.201-4.122-3.181-7.134l.08-9.985c.07-2.781.43-4.723 3.432-6.364l10.796-5.873c1-.52 2.4-.58 3.381-.58M15.15 9.054L33.2 19.66M7.144 15.237l15.869 9.135c.67.35 1.34.23 2.001-.06l15.779-8.915m-8.085 4.043c.6.34.49 1.17.49 1.93v4.623m-9.235-1.27v17.59"/>`),
		g.Group(children),
	)
}