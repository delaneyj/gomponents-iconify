package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ipfslite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Zm-30.35-24v11m12.77-11h5.5m-5.5 5.5h3.57m-3.57-5.5v11"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.92 29.5v-11h3.6a3.7 3.7 0 0 1 0 7.39h-3.6m17.86 2.4a3.07 3.07 0 0 0 2.69 1.21h1.63a2.75 2.75 0 0 0 2.75-2.75h0A2.75 2.75 0 0 0 35.1 24h-1.8a2.75 2.75 0 0 1-2.74-2.75h0a2.75 2.75 0 0 1 2.74-2.75h1.63a3.08 3.08 0 0 1 2.7 1.21"/>`),
		g.Group(children),
	)
}