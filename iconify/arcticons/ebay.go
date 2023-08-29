package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ebay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M38.5 24.81v3.357a2.478 2.478 0 0 1-2.478 2.478h-.02h0a2.448 2.448 0 0 1-1.76-.729"/><path d="M38.5 20.702v4.107a2.488 2.488 0 0 1-2.488 2.489h-.01h0a2.488 2.488 0 0 1-2.479-2.489v-4.107m-19.356 5.297a2.498 2.498 0 0 1-2.149 1.299h0a2.488 2.488 0 0 1-2.498-2.479V23.19a2.488 2.488 0 0 1 2.488-2.488h.01a2.488 2.488 0 0 1 2.479 2.488V24H9.5m20.995.81a2.488 2.488 0 0 1-2.488 2.488h0a2.488 2.488 0 0 1-2.478-2.489V23.19a2.488 2.488 0 0 1 2.478-2.488h0a2.488 2.488 0 0 1 2.488 2.488m0 4.108v-6.596M17.524 23.19a2.488 2.488 0 0 1 2.489-2.488h0a2.488 2.488 0 0 1 2.478 2.488v1.62a2.488 2.488 0 0 1-2.478 2.488h0a2.488 2.488 0 0 1-2.489-2.489m0 2.489v-9.943"/></g><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/>`),
		g.Group(children),
	)
}