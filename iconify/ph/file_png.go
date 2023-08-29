package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePng(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M60 144H44a8 8 0 0 0-8 8v56a8 8 0 0 0 16 0v-8h8a28 28 0 0 0 0-56Zm0 40h-8v-24h8a12 12 0 0 1 0 24Zm164 16.87a8 8 0 0 1-2.22 5.53A30.06 30.06 0 0 1 200 216c-17.65 0-32-16.15-32-36s14.35-36 32-36a29.45 29.45 0 0 1 16.48 5.11a8 8 0 0 1-9 13.27A13.21 13.21 0 0 0 200 160c-8.82 0-16 9-16 20s7.18 20 16 20a13.57 13.57 0 0 0 8-2.72V192a8 8 0 0 1 0-16h8a8 8 0 0 1 8 8ZM156 152v56a8 8 0 0 1-5.56 7.62a7.91 7.91 0 0 1-2.44.38a8 8 0 0 1-6.51-3.35L116 177v31a8 8 0 0 1-16 0v-56a8 8 0 0 1 14.51-4.65L140 183v-31a8 8 0 0 1 16 0ZM48 120a8 8 0 0 0 8-8V40h88v48a8 8 0 0 0 8 8h48v16a8 8 0 0 0 16 0V88a8 8 0 0 0-2.34-5.66l-56-56A8 8 0 0 0 152 24H56a16 16 0 0 0-16 16v72a8 8 0 0 0 8 8Zm112-68.69L188.69 80H160Z"/>`),
		g.Group(children),
	)
}