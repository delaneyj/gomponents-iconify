package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Npm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M0 9.781v10.667h8.885v1.771H16v-1.771h16V9.781zm8.885 8.88H7.114v-5.333H5.333v5.333H1.781v-7.104h7.104zm5.334 0v1.787h-3.552v-8.891h7.115v7.109h-3.563zm16.005 0h-1.776v-5.333h-1.781v5.333h-1.781v-5.333h-1.771v5.333h-3.563v-7.104h10.672zm-16.005-5.328H16v3.557h-1.781z"/>`),
		g.Group(children),
	)
}