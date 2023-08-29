package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Avocado(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M386 1028q-81 0-150.5-21T113 946.5T30 844T0 702q0-60 17-127t38.5-118T109 329t51-134q28-82 88.5-138.5T385 0t136.5 56.5T610 195q19 56 51.5 134t54 129.5T754 577t17 126q0 104-52.5 179T580 992.5T386 1028zm192-739q-11-36-30-73t-43.5-72t-56-57.5t-63-22.5t-63 22t-56.5 57t-43.5 72.5T193 289q-20 61-55 145.5t-54.5 142T64 679q0 91 44 156.5t115.5 97t162 31.5t162-31t115-96.5T707 679q0-45-19-102t-54.5-142.5T578 289zM385.5 899q-79.5 0-136-56.5T193 706q0-54 25.5-113t71-101.5t96-42.5t96.5 42.5T553 593t25 113q0 80-56.5 136.5t-136 56.5z"/>`),
		g.Group(children),
	)
}