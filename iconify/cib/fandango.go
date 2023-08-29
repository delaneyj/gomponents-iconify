package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fandango(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m18.219 9.276l-7.484 2.052l1.52 5.631l7.485-2.052l1.525 5.624l-7.484 1.991l1.459 5.629l-7.491 1.989l-1.853-6.755c1.125-1.323 1.588-3.183 1.125-4.969c-.532-1.791-1.855-3.181-3.443-3.776l-1.86-6.957l15.041-3.975zm9.474-.265L25.245.068L0 6.755l2.453 9.011c1.453.271 2.781 1.391 3.177 2.984c.463 1.589-.136 3.177-1.261 4.24l2.385 8.943l25.245-6.755l-2.385-8.943c-1.527-.271-2.781-1.391-3.245-2.984c-.4-1.589.131-3.245 1.324-4.24z"/>`),
		g.Group(children),
	)
}