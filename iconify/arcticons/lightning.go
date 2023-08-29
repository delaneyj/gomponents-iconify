package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lightning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.53 7.09a7.28 7.28 0 0 1 6.78 4.63A5.74 5.74 0 0 1 35 15.89a4.89 4.89 0 0 1 0 .54a5.19 5.19 0 0 1 7.52 4.64h0c0 2.29-1.21 3.77-3.39 5.06l-.3.13H25.19l-.74 3.48l5.81.47l-7 13.45l2.14-10.92l-5.4-.3l1.3-6.18H9L8.41 26c-2.12-1.29-2.91-2.66-2.91-4.79a5.06 5.06 0 0 1 5.06-5.05a5.2 5.2 0 0 1 1 .09a7.55 7.55 0 0 1-.26-1.89a7.27 7.27 0 0 1 7.27-7.27Zm2.8 19.17h3.86"/>`),
		g.Group(children),
	)
}