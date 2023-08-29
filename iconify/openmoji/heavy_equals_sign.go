package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeavyEqualsSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M54.939 32.334H17.061v-9.905h37.871m.007 26.977H17.061v-9.905h37.871"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.061 22.352h37.878v10.296H17.061zm0 17h37.878v10.296H17.061z"/>`),
		g.Group(children),
	)
}