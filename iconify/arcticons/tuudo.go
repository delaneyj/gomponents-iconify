package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tuudo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="4.551" height="6.029" x="33.95" y="21.521" rx="2.275" ry="2.275"/><path d="M13.852 21.521v3.754a2.275 2.275 0 0 0 2.275 2.276h0a2.275 2.275 0 0 0 2.276-2.276v-3.754m0 3.754v2.275m2.139-6.029v3.754a2.275 2.275 0 0 0 2.275 2.276h0a2.275 2.275 0 0 0 2.275-2.276v-3.754m0 3.754v2.275m-14.398-7.906v7.906M9.5 21.521h2.389m19.857 2.275a2.275 2.275 0 0 0-2.275-2.275h0a2.275 2.275 0 0 0-2.276 2.275v1.48a2.275 2.275 0 0 0 2.276 2.274h0a2.275 2.275 0 0 0 2.275-2.275m0 2.275v-9.1"/></g>`),
		g.Group(children),
	)
}