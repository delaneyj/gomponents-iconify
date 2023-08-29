package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Safetypin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M951.816 951q-73 73-175.5 73t-175.5-73l-527-527q-73-73-73-175.5t73-175.5t175.5-73t175.5 73l527 527q73 73 73 175.5t-73 175.5zm-439-695q-26 0-70-1t-60-1t-38 4.5t-37 15.5q-21-29-21-70.5t-12-51.5q-25-25-60.5-25t-61 25.5t-25.5 61t25 60.5q12 11 31 11t43.5 2.5t46.5 19.5q-10 15-15 36t-5 37.5t1.5 59.5t1.5 72l256 264v-8q0-106 75-181t181-75h8zm256.5 320q-79.5 0-136 56.5t-56.5 136t56.5 135.5t136 56t135.5-56t56-135.5t-56-136t-135.5-56.5z"/>`),
		g.Group(children),
	)
}