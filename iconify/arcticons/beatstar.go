package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beatstar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.117 16.149H8.749L4.5 30.373h39l-4.249-14.224h-2.368"/><path d="M38.765 25.041v2.744H9.235v-2.744l2.367-10.673h24.796l2.367 10.673zM11 25.389h26"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 30.373h39v3.259h-39z"/>`),
		g.Group(children),
	)
}