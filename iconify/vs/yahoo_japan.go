package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YahooJapan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="m1733 481l46-73h-715l12 78q19 0 103 11t118 18q-3 30-85.5 103.5t-174 146.5T934 855q-22-32-160-211T568 366q26-7 135.5-15T829 340l12-71H9l-9 75q18 4 128.5 18.5T253 382q66 50 262.5 293.5T726 959v256q0 29-20 45.5t-56 26.5q-86 24-163 17l-21 85q59 2 114 2t145.5-1.5T837 1388q400 0 421 3l10-96l-247-17q-5-298 0-324q13-37 131-142.5t237-196t143-96.5q185-35 201-38zm464 6l-390 663l-135-38l217-718zm-674 896l174 49l67-149l-179-50z"/>`),
		g.Group(children),
	)
}