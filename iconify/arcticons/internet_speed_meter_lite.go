package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InternetSpeedMeterLite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-8.336 27.446V15.148"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m38 22.964l-5.835-7.91l-5.836 7.91m-10.494-7.91v17.798m5.836-7.816l-5.836 7.91L10 25.036"/>`),
		g.Group(children),
	)
}