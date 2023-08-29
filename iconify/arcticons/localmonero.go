package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Localmonero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.383 28.418h4.322v-8.836L24 26.407l6.29-6.825v8.836h4.327"/><path d="M35.5 24c0 6.351-5.149 11.5-11.5 11.5S12.5 30.351 12.5 24S17.65 12.5 24 12.5S35.5 17.65 35.5 24Z"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 24L13.75 6.246h20.5L44.5 24L34.25 41.754h-20.5L3.5 24Z"/>`),
		g.Group(children),
	)
}