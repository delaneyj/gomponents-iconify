package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Theater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTheater0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" fill-rule="evenodd" stroke-linejoin="round" d="M8 6h32a2 2 0 0 1 2 2v28L32 22.005c-2.667 1.585-5.333 2.378-8 2.378s-5.333-.793-8-2.378L6 36V8a2 2 0 0 1 2-2Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M6 42h36"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTheater0)"/>`),
		g.Group(children),
	)
}