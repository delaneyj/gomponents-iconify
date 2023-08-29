package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hiit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.053 45.107a19.163 19.163 0 0 1-6.142-30.19m2.443-2.279A19.163 19.163 0 0 1 39.15 8.812"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.911 14.917l-1.923-1.91l-1.407 1.323l-1.923-1.969l5.404-4.949l1.816 1.93l-1.475 1.34l1.951 1.956m10.024-.648l4.33 16.106l-8.934 10.223m8.935-10.223l5.745 2.557"/>`),
		g.Group(children),
	)
}