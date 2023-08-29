package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cocktail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M766.5 99q-4.5 16-20 24t-32.5 4L567 84l-28 108h5q13 0 22.5 9.5T576 224v736q0 27-19 45.5t-45 18.5H256q-27 0-45.5-19T192 960V384q-80 0-136-56T0 192T56 56T192 0t136 56t56 136h73l42-160q4-12 11-18.5t18-9t17-3t21-.5l170 50q18 4 26.5 18t4 30zM256 302q29-17 46-46h-46v46zM192 64q-53 0-90.5 37.5T64 192t37.5 90.5T192 320v-96q0-13 9.5-22.5T224 192h96q0-53-37.5-90.5T192 64zm180 192q-14 41-44.5 71.5T256 372v76h133l51-192h-68zm140 38l-41 154h41V294z"/>`),
		g.Group(children),
	)
}