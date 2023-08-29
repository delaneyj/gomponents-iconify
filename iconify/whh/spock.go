package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m765 213l-60 491q-6 54-18 106.5T662 900t-25 66t-20 44t-8 14H257q-5-5-13.5-14t-31-34.5t-39-48.5t-30.5-50t-14-45q0-38-27.5-122T46 561l-27-65Q0 477 0 449.5t19-47T65 383t46 20q0 2 1.5 6t5 15.5t7.5 23t9.5 27t11 29t12 27.5t12.5 23.5t12 15.5t11 6l-64-437q-4-26 10.5-48T178 65q17-1 34.5 0t31 17t13.5 46V64q0-27 19-45.5T321.5 0t45 18.5T385 64l64 384q7-2 15-10.5t11-21.5l50-297q4-27 23.5-42.5t44-11.5T631 91t10 48v53q23-64 64-64q37 0 53.5 28.5T765 213z"/>`),
		g.Group(children),
	)
}