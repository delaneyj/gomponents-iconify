package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mlauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m28.035 37.69l.854-10.567l-2.547 3.434l-1.41-.023l-1.84-5.372l-.862 11.178m10.136 1.393l3.028-17.408"/><path d="m35.017 24.623l1.589.587l4.488-12.728l-.876-3.226l-10.19 6.61l-1.033 6.866l-2.913 4.186l-2.74-8.763l.383-2.852L15.767 5.5l-.912 15.855c1.896.228 1.405-1.358 1.605-2.946l-1.113 15.682m5.38 1.649l2.849 1.512l-2.673 3.373l-13.997-5.413l1.355-1.284c2.931.18 6.113.5 8.772.53m12.202 3.61l-2.548.096l.373 4.336c2.847-1.442 5.763-2.716 9.176-2.791l-.06-1.5l-4.634-.34"/></g>`),
		g.Group(children),
	)
}