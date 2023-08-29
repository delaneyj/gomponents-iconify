package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WipePhone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M534 1664q-20 52-20 107v10q0 5 1 11H384v-128h150zm965 256h293v128H837l-147-148q-24-25-37-57t-13-67q0-35 13-67t38-58l813-814l538 539l-543 544zm5-902l-389 390l357 358l389-390l-357-358zm-187 902l65-64l-358-357l-242 242q-14 14-14 35t14 35l108 109h427zm-767 0q21 41 47 68t60 60H128q-27 0-50-10t-40-27t-28-41t-10-50V128q0-27 10-50t27-40t41-28t50-10h1024q27 0 50 10t40 27t28 41t10 50v752l-128 128V128H128v1792h422z"/>`),
		g.Group(children),
	)
}