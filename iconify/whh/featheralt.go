package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Featheralt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M532.56 410q-66-66-165.5-149.5T201.56 127l-67-51q-12-12-29-12t-29 12q-3 3-4 7t.5 9t4 8.5t7.5 9t8 9t8.5 8.5t7.5 7q25 18 65.5 49t139.5 114t167 151q42 42 111.5 111t119.5 119t89 90q58 61 111 118t77 83.5l24 26.5q12 12 12 13t-12 13q-10 11-71.5-32.5t-82.5-64.5q-60-61-162-164q-71-18-175.5-76t-195.5-122.5t-158-116.5t-85-70q-26-26-51-91.5t-30.5-131T26.56 51q47-48 132-50.5t154 37.5q76 43 172 131.5t176 187t130 196.5t40 156q-147-148-298-299z"/>`),
		g.Group(children),
	)
}