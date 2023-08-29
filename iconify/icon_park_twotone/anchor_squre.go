package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnchorSqure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAnchorSqure0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 38c0-9.941-8.059-20-18-20S6 28.059 6 38m14-24H10m28 0H28"/><circle cx="24" cy="14" r="4" fill="#555"/><path fill="#555" d="M20 10h8v8h-8zm18 1h6v6h-6zM4 11h6v6H4z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAnchorSqure0)"/>`),
		g.Group(children),
	)
}