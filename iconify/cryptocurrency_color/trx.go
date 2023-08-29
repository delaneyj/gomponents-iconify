package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#EF0027"/><path fill="#FFF" d="M21.932 9.913L7.5 7.257l7.595 19.112l10.583-12.894l-3.746-3.562zm-.232 1.17l2.208 2.099l-6.038 1.093l3.83-3.192zm-5.142 2.973l-6.364-5.278l10.402 1.914l-4.038 3.364zm-.453.934l-1.038 8.58L9.472 9.487l6.633 5.502zm.96.455l6.687-1.21l-7.67 9.343l.983-8.133z"/></g>`),
		g.Group(children),
	)
}