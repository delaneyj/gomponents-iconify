package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Meualelo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.213 26.874a2.059 2.059 0 0 1-2.053-2.053v-1.334a2.053 2.053 0 0 1 4.106 0v1.334a2.059 2.059 0 0 1-2.053 2.053Zm-9.955-1.027a1.986 1.986 0 0 1-1.745 1.027a2.059 2.059 0 0 1-2.053-2.053v-1.334a2.053 2.053 0 0 1 4.106 0v.718H22.46m7.698-5.542v7.184a.97.97 0 0 0 1.026 1.027h.308m-13.034-8.211v7.184a.97.97 0 0 0 1.026 1.027h.308m-4.926-2.053a2.053 2.053 0 0 1-4.105 0v-1.334a2.053 2.053 0 0 1 4.105 0m0-2.053v5.44"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.895 32.21a6.392 6.392 0 0 0-2.999.47a9.237 9.237 0 1 1 0-17.36a6.392 6.392 0 0 0 2.999.47c1.438-.212 5.131-2.456 5.131-2.456a12.316 12.316 0 1 1 0 21.331s-3.693-2.244-5.131-2.454Z"/>`),
		g.Group(children),
	)
}