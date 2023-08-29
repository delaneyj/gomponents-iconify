package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Imepay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.37 20.24l3 2.85a.52.52 0 0 0 .83-.12l7.09-12.66L16.16 14a.53.53 0 0 0-.25.9l4 3.87l7.67-5.11l-7.67 5.11l1.49 1.47l-2.4 1.93a.42.42 0 0 1-.68-.32v-4.57m16.43 14.49v3.72A2.75 2.75 0 0 1 32 38.24h0a2.73 2.73 0 0 1-1.94-.81"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.75 27.24v4.53A2.75 2.75 0 0 1 32 34.52h0a2.75 2.75 0 0 1-2.75-2.75v-4.53m-2.5 4.53A2.75 2.75 0 0 1 24 34.52h0a2.75 2.75 0 0 1-2.75-2.75V30A2.75 2.75 0 0 1 24 27.24h0A2.75 2.75 0 0 1 26.75 30m0 4.52v-7.28m-13.5 4.53A2.75 2.75 0 0 0 16 34.52h0a2.75 2.75 0 0 0 2.75-2.75V30A2.75 2.75 0 0 0 16 27.24h0A2.75 2.75 0 0 0 13.25 30m0-2.76v11"/>`),
		g.Group(children),
	)
}