package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amazon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M379 297q6-2 9 2.5t-2 8.5q-34 25-81 39t-92 14Q91 361 2 280q-3-3-1-5.5t6-.5q96 56 211 56q83 0 161-33zm46-26q5 6-2.5 31.5T399 342q-3 3-5 2t-1-5q18-45 12-53t-54-2q-4 0-4.5-2t2.5-4q18-13 46-13.5t30 6.5zM237 113v-6q0-22-6-30q-7-11-23-11q-28 0-33 25q-2 8-8 8l-40-4q-8-2-6-9q6-34 32.5-49T214 22q41 0 63 21q3 3 5.5 6t4.5 7.5t3.5 7t2 8t1.5 8t1 9V180q0 17 16 38q5 7 0 12q-16 12-32 27q-5 4-10 1q-11-9-24-28q-17 18-32 24.5t-37 6.5q-27 0-44.5-17T114 196q0-49 44-69q17-7 79-14zm-8 87q8-14 8-45v-9q-62 0-62 42q0 14 6.5 22.5T200 219q18 0 29-19z"/>`),
		g.Group(children),
	)
}