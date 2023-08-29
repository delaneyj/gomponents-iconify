package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ATextWidget(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.8 27.4h2.1m-2.1-4.3h2.1m-2.1 2.2h1.3m-1.3-2.2v4.3m3.6-4.3l2.8 4.3m0-4.3l-2.8 4.3m-8-4.3h2.8m-1.4 4.3v-4.3M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Zm-3.7 17.6h2.8m-1.4 4.3v-4.3"/>`),
		g.Group(children),
	)
}