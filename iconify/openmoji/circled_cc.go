package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircledCc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="26.68" fill="#fff" fill-rule="evenodd"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round"><path stroke-miterlimit="10" stroke-width="5.84" d="M52.34 44.08a7.302 7.302 0 0 1-12.146-5.464v-5.84v5.84v-5.84a7.302 7.302 0 0 1 12.146-5.464M31.78 44.69a7.302 7.302 0 0 1-12.146-5.464v-5.84v5.84v-5.84a7.302 7.302 0 0 1 12.146-5.464" clip-rule="evenodd"/><circle cx="36" cy="36" r="26.68" stroke-width="4.74"/></g>`),
		g.Group(children),
	)
}