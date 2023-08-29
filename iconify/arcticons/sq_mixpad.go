package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SqMixpad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.5 32.116l-5.364-2.868m2.883-1.993a3.024 3.024 0 0 0 .924-2.08v-4.521c0-2.513-3.718-4.524-8.363-4.524c-3.477 0-6.26 1.126-7.44 2.747m-.611 3.214v3.084c0 2.513 3.715 4.524 8.05 4.524c1.902 0 2.98.114 4.38-.456M4.83 28.003a11.902 11.902 0 0 0 6.603 1.536h3.962c3.632 0 6.603-1.536 6.603-3.413s-2.971-3.414-6.603-3.414h-4.292c-3.632 0-6.603-1.536-6.603-3.414s2.971-3.414 6.603-3.414h3.962c2.971 0 4.952.342 6.603 1.537"/>`),
		g.Group(children),
	)
}