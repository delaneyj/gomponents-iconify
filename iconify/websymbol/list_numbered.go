package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListNumbered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 118v213H331V118h669zm0 334v213H331V452h669zM196 139v192h-42V206h-46v-29q49 0 56-38h32zm804 648v213H331V787h669zM232 626v33H91q1-28 17-48.5t34.5-30t34-23.5t15.5-30q0-12-7.5-19.5T164 500q-29 0-31 41H96q0-34 18-54.5t52-20.5q28 0 46.5 16t18.5 43q0 25-16.5 41t-40 30.5T141 626h91zm1 303q0 29-21 45.5T162 991q-32 0-51.5-17.5T91 925v-3h38q0 37 32 37q14 0 23-8t9-22t-10-21.5t-25-6.5h-9v-27h6q34 0 34-24q0-11-7.5-17.5T162 826q-29 0-29 33H96q1-30 18.5-47t47.5-17q26 0 45 13.5t19 38.5q0 13-7.5 24T199 885q34 8 34 44z"/>`),
		g.Group(children),
	)
}