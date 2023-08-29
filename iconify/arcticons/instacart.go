package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instacart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.56 3.55c1.08 0 6 .45 6 4.41s-6 8.58-7.5 8.58S28 11.48 28 9s.3-5.45 4.56-5.45Zm.1 25.09C24.23 37 8.7 43.76 7.31 42.37S12.62 25.46 21.05 17c1.3-1.3 5-2.93 9.76 1.82s3.15 8.51 1.85 9.8Zm6-7.23c-1.81 0-5.57-1.22-5.57-2.32s3.43-5.57 6.37-5.57s3.27 3.66 3.27 4.46c0 3.19-2.27 3.44-4.07 3.44Z"/>`),
		g.Group(children),
	)
}