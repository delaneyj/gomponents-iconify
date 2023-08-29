package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Insular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.47 7.88H8.53L4.5 34.53h39L39.47 7.88zm-30.94 0L43.5 34.53M14.61 7.88l28.1 21.41M20.7 7.88l21.22 16.17M26.78 7.88l14.34 10.93M32.87 7.88l7.46 5.69M7.9 12.04l29.52 22.49M7.27 16.2l24.06 18.33M6.64 20.35l18.61 14.18M6.02 24.51l13.14 10.02M5.39 28.67l7.69 5.86m28.44 5.59H6.48L4.5 34.53h39l-1.98 5.59z"/>`),
		g.Group(children),
	)
}