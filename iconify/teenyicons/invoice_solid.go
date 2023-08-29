package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InvoiceSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M10 7.995V9H8V7.995h2ZM10 6v.995H8V6h2ZM7 6H5v.995h2V6Zm0 1.995H5V9h2V7.995Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM4 4h3V3H4v1Zm7 1H4v5h7V5Zm0 7H8v-1h3v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}