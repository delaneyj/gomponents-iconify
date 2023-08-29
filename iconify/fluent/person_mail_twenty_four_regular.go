package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonMailTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 15.5c0-.563.186-1.082.5-1.5H5.253a2.249 2.249 0 0 0-2.25 2.25v.577c0 .892.32 1.756.9 2.435c1.565 1.834 3.951 2.74 7.097 2.74h.05A2.514 2.514 0 0 1 11 21.5v-.999c-2.738 0-4.704-.745-5.957-2.213a2.25 2.25 0 0 1-.54-1.461v-.578a.75.75 0 0 1 .75-.749H11Zm0-13.495a5 5 0 1 1 0 10a5 5 0 0 1 0-10Zm0 1.5a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7Zm6.51 15.422l-5.49-3.203A2 2 0 0 1 14 14h7a2 2 0 0 1 2 1.97l-5.49 2.957Zm.227 1.014l5.264-2.834V21a2 2 0 0 1-2 2h-7a2 2 0 0 1-2-2v-4.13l5.248 3.062a.5.5 0 0 0 .489.009Z"/>`),
		g.Group(children),
	)
}