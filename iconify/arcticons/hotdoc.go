package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hotdoc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.906 40.455L24.1 45.5l-6.852-4.981l-8.38-1.245l-1.323-8.368L2.5 24.1l4.981-6.852l1.245-8.38l8.368-1.323L23.9 2.5l6.852 4.981l8.38 1.245l1.323 8.368L45.5 23.9l-4.981 6.852l-1.245 8.38Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.783 19.91v-7.954h8.328v8.141h8.277v7.938h-8.21v8.01H20.01v-8.11h-8.4v-8.033Z"/>`),
		g.Group(children),
	)
}