package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileDe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Zm-24.6.3v36.4m4.8-23.8v11.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.7 22.629a4.225 4.225 0 0 1 4.225-4.225h0a4.225 4.225 0 0 1 4.225 4.225V29.6M20.7 18.404V29.6m8.45-6.971a4.225 4.225 0 0 1 4.225-4.225h0a4.225 4.225 0 0 1 4.225 4.225V29.6"/>`),
		g.Group(children),
	)
}