package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whiteboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M128 64q0-26 19-45t45-19h1408q26 0 45 19t19 45v1024h98q12 0 21 9t9 21v68q0 12-9 21t-21 9h-579l131 487q7 25-6.5 48.5T1268 1782t-48.5-6.5t-29.5-38.5l-140-521h-90v384q0 26-19 45t-45 19t-45-19t-19-45v-384h-90l-140 521q-7 25-29.5 38.5T524 1782t-39.5-30.5t-6.5-48.5l131-487H30q-12 0-21-9t-9-21v-68q0-12 9-21t21-9h98V64zm128 1024h1280V128H256v960zm128-64V915q0-14 9.5-24t23.5-10h318q14 0 23.5 10t9.5 24v109H384z"/>`),
		g.Group(children),
	)
}