package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M500 2000q-253 2-394-26q-48-10-77-64T0 1784v-399q0-151 42-291t99-193q60-56 124.5-159T343 592v-20h318v20q13 47 77.5 150T863 901q56 53 96.5 192.5T1000 1385v399q0 73-27.5 127t-74.5 63q-141 28-398 26zM394 0h217q17 0 26 3t16.5 18.5T661 66v386H343V66q0-29 7.5-44.5t17-18.5T394 0z"/>`),
		g.Group(children),
	)
}