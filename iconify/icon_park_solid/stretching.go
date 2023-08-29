package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stretching(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSStretching0"><g fill="none"><g stroke="#fff" stroke-width="4" clip-path="url(#ipSStretching1)"><path stroke-linecap="round" stroke-linejoin="round" d="M23 6H8a2 2 0 0 0-2 2v32a2 2 0 0 0 2 2h32a2 2 0 0 0 2-2V25"/><path stroke-linecap="round" d="M24.001 16v8M42 6v8m-9.999 10h-8"/><path d="M42 6L24 24"/><path stroke-linecap="round" d="M42 6h-8"/></g><defs><clipPath id="ipSStretching1"><path fill="#000" d="M0 0h48v48H0z"/></clipPath></defs></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSStretching0)"/>`),
		g.Group(children),
	)
}