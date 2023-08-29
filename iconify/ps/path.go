package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Path(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M116 23Q22 61 5 153q-11 62 12 113q6 11 14 23q2 3 8 9q10-4 15-7q9-5 17.5-12T90 262.5t16-14.5q-39-51-10-102t90-65q56-14 114 9.5t74 75.5t-22 85q-19 17-46 22q-14 3-31 1q-5 0-17-2q-2-1-5-1.5t-5-1t-3.5-2t-2.5-3.5t-1-5V140q0-6-.5-6.5t-5.5-.5q-12-2-16-2q-39-2-57 1q-8 5-8 7v63q0 18 1.5 64t.5 70q-4 46-16 52q-10 4-27.5-3.5T92 376q0 5-1.5 25.5t1 34T106 454q42 17 76 5q37-13 53-48.5t10-75.5q61 18 119-8.5t85-81.5q18-39 11.5-83.5T426 84q-34-40-89-61T224 2.5T116 23z"/>`),
		g.Group(children),
	)
}