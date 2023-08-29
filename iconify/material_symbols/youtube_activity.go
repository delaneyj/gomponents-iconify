package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YoutubeActivity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.4 17.95q-.675.025-1.288.038T13.026 18H12q-1.775 0-3.325-.05q-1.325-.05-2.613-.138T4.2 17.576q-.65-.175-1.125-.65t-.65-1.125q-.15-.575-.237-1.4t-.138-1.575Q2 11.925 2 11t.05-1.825q.05-.75.138-1.575t.237-1.4q.175-.65.65-1.125t1.125-.65q.575-.15 1.863-.238t2.612-.137Q10.225 4 12 4t3.325.05q1.325.05 2.613.138t1.862.237q.65.175 1.125.65t.65 1.125q.15.575.238 1.4t.137 1.575q.05.9.05 1.825v.425q-.475-.2-.975-.313T20 11q-2.075 0-3.538 1.463T15 16q0 .525.1 1.012t.3.938ZM10 14l5.2-3L10 8v6Zm9 5v-2h-2v-2h2v-2h2v2h2v2h-2v2h-2Z"/>`),
		g.Group(children),
	)
}