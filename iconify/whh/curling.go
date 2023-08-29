package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Curling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 864q0 28-34.5 53T765 960H259q-62-18-96.5-43T128 864L0 800V384l128-64q0-28 34.5-53t96.5-43h320l-96-96H192q-27 0-45.5-19T128 63.5t18.5-45T192 0h320q1 0 3 .5l2 .5q7 0 14 2t15 7q2 1 3 2q10 7 11 7l205 205q62 18 96.5 43t34.5 53l128 64v416zm-704-32q0 37 109 64h422q109-27 109-64l128-68V510q-67 30-186.5 48T512 576t-261.5-18T64 510v254zm656-448h-4q-51 29-139 46.5T512 448t-193-17.5T180 384h-4L64 434v12q66 30 185.5 48T512 512t262.5-18T960 446v-12z"/>`),
		g.Group(children),
	)
}