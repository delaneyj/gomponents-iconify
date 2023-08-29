package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fedex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.647 24h5.057m-5.057-5.153h5.057M27.647 24h3.287m-3.287-5.153v10.306m0 0h5.057M8.5 18.847h5.057M8.5 24h3.287M8.5 18.847v10.306m31-6.827l-5.057 6.827m5.057 0l-5.057-6.827m-16.208 5.531a2.44 2.44 0 0 1-2.149 1.288a2.56 2.56 0 0 1-2.529-2.576v-1.675a2.53 2.53 0 1 1 5.058 0v.902h-5.058m12.129-.894a2.53 2.53 0 1 0-5.058 0v1.675a2.53 2.53 0 1 0 5.058 0m0 2.576V18.847"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.461 6.675a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33.04a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}