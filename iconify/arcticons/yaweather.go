package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yaweather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 25.06a14.8 14.8 0 0 1 6.09 1.52a1.13 1.13 0 0 0 1.1 0a12.6 12.6 0 0 1 8.54-1.33a1.15 1.15 0 0 0 1.36-1.13a17 17 0 0 0-5-12.06h0a17 17 0 0 0-10.95-4.95V5.64a1.14 1.14 0 0 0-2.28 0v1.47A17.08 17.08 0 0 0 6.91 24a1.16 1.16 0 0 0 .39.87a1.14 1.14 0 0 0 .92.27a12.51 12.51 0 0 1 1.83-.14a12.65 12.65 0 0 1 6.4 1.75a1.19 1.19 0 0 0 .58.15a1.16 1.16 0 0 0 .57-.15a13.72 13.72 0 0 1 6.4-1.7v15a3.44 3.44 0 0 0 3.43 3.43h0a3.44 3.44 0 0 0 3.43-3.43v-1.2M17.3 26.7a43.3 43.3 0 0 1 2.26-14.2c1.26-3.38 2.92-5.39 4.44-5.39h0c1.52 0 3.18 2 4.44 5.39h0a43.24 43.24 0 0 1 2.26 14.19M22.86 7.11h2.28"/>`),
		g.Group(children),
	)
}