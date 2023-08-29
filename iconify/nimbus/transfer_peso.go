package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransferPeso(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.75 2.5H1.25A1.2 1.2 0 0 0 0 3.64v8.72a1.2 1.2 0 0 0 1.25 1.14h13.5A1.2 1.2 0 0 0 16 12.36V3.64a1.2 1.2 0 0 0-1.25-1.14zm0 9.75H1.25v-8.5h13.5z"/><path fill="currentColor" d="M7 8.62h2a.34.34 0 0 1 .33.38a.33.33 0 0 1-.33.29H7.08A.33.33 0 0 1 6.75 9H5.49a1.58 1.58 0 0 0 1.58 1.54h.31v1.26h1.24v-1.26H9A1.58 1.58 0 0 0 10.56 9a1.51 1.51 0 0 0-.34-1A1.59 1.59 0 0 0 9 7.38H7A.34.34 0 0 1 6.69 7a.13.13 0 0 1 .01 0a.34.34 0 0 1 .3-.3h1.94a.34.34 0 0 1 .33.3h1.25a1.59 1.59 0 0 0-1.58-1.55h-.32V4.2H7.37v1.25H7A1.6 1.6 0 0 0 5.44 7a1.55 1.55 0 0 0 .35 1A1.59 1.59 0 0 0 7 8.62z"/>`),
		g.Group(children),
	)
}