package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XingSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M685 637q0-1-126-222q-21-34-52-34H323q-18 0-26 11q-7 12 1 29l125 216v1L227 984q-9 14 0 28q8 13 24 13h185q31 0 50-36zm624-497q-7-12-24-12h-187q-30 0-49 35L638 892q1 2 262 481q20 35 52 35h184q18 0 25-12q8-13-1-28L900 892v-1l409-723q8-16 0-28zm227 148v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288z"/>`),
		g.Group(children),
	)
}