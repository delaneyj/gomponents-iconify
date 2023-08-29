package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="m806 1199l58-58l-149-149l-58 58v68h81v81h68zm331-588q0-7-4-10q-3-4-10-4q-3 0-5 1q-4 2-5 3L769 945q-1 1-3 5q-1 2-1 5q0 7 4 11q3 3 10 3h5q4-2 6-4l343-343q3-3 3-5q1-2 1-6zm-34-122l264 264l-527 527H576v-264zm433 61q0 16-6 31t-17 26l-106 105l-263-263l105-105q8-9 26-18q15-6 31-6t31 6q18 9 27 18l149 148q11 13 17 27q6 15 6 31zM128 1408h1024v-320l128-128v480q0 25-13 49q-14 21-35 34q-20 13-48 13H96q-28 0-48-13q-21-13-35-34q-11-22-13-49V96q2-28 13-48q13-22 35-35Q71 0 96 0h640q29 0 58 10q29 8 59 24q24 12 47 34l251 253l-93 93l-249-256q-10-11-37-20q-34-10-52-10H128v1280z"/>`),
		g.Group(children),
	)
}