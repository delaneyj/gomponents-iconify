package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1472 736v128q0 53-32.5 90.5T1355 992H651l293 294q38 36 38 90t-38 90l-75 76q-37 37-90 37q-52 0-91-37L37 890Q0 853 0 800q0-52 37-91L688 59q38-38 91-38q52 0 90 38l75 74q38 38 38 91t-38 91L651 608h704q52 0 84.5 37.5T1472 736z"/>`),
		g.Group(children),
	)
}