package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magento(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 32l-5.937-3.427v-16l3.959-2.287v16l1.979 1.203l1.979-1.203v-16l3.964 2.287v16zM29.855 8v16l-3.959 2.287v-16L16 4.574l-9.901 5.713v16L2.146 24V8l13.855-8z"/>`),
		g.Group(children),
	)
}