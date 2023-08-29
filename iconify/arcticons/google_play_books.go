package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlayBooks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.055 27.738c1.35-.727 2.284-2.18 2.284-3.738c-.104-1.557-1.038-3.01-2.388-3.738h0L15.202 5.623C14.58 5.208 13.75 5 12.918 5C10.945 5 9.18 6.35 8.765 8.115h0c-.104.311-.104.727-.104 1.142v29.59c0 .415 0 .727.104 1.142v-.104h0C9.285 41.65 10.945 43 12.918 43c.83 0 1.557-.208 2.18-.623h0l25.957-14.639Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.202 5.623v17.222a.5.5 0 0 0 .854.354l4.95-4.95l4.93 4.93a.5.5 0 0 0 .854-.353v-10.4"/>`),
		g.Group(children),
	)
}