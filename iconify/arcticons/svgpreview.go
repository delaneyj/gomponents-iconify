package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Svgpreview(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 25.873V12.864l2.828-2.828h23.646L42.5 20.563v8.06H8.25l-2.75-2.75zm5.327 7.225v-2.389H37.64v2.754m-6.002 5.852H16.397"/><circle cx="13.317" cy="36.47" r="4.196" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.717" cy="36.47" r="4.196" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.397 28.623v2.086m15.909-2.086v2.086m-5.358-20.673l-.655-2.702h-1.62l-.656 2.702v10.527H42.5"/>`),
		g.Group(children),
	)
}