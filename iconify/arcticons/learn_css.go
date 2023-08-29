package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LearnCss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.623 12.452a3.876 3.876 0 0 0 2.623.656h.656a2.131 2.131 0 0 0 0-4.263h-1.476a2.17 2.17 0 0 1-2.131-2.132a2.17 2.17 0 0 1 2.131-2.131h.656c1.476 0 2.131.164 2.623.656m4.673 7.214a3.876 3.876 0 0 0 2.623.656h.656a2.131 2.131 0 0 0 0-4.263h-1.475a2.17 2.17 0 0 1-2.132-2.132a2.17 2.17 0 0 1 2.131-2.131h.656c1.476 0 2.132.164 2.624.656M17.278 11.55a3.173 3.173 0 0 1-2.787 1.64a3.289 3.289 0 0 1-3.28-3.28V7.78a3.289 3.289 0 0 1 3.28-3.28a3.173 3.173 0 0 1 2.787 1.64m-3.886 12.39h21.055l-2.171 22.943l-8.288 2.027l-8.287-2.027l-.575-4.72m-.724-7.729h18.763"/>`),
		g.Group(children),
	)
}