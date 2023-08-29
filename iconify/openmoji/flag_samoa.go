package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagSamoa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><path fill="#1e50a0" d="M5 17h32v21H5z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m20.081 24.089l1.232-4l1.24 3.998l-3.238-2.469l4-.003l-3.234 2.474zm0 11.882l1.232-4l1.24 3.998l-3.238-2.469l4-.003l-3.234 2.474zm6.5-8.5l1.232-4l1.24 3.998L25.815 25l4-.003l-3.234 2.474zm-13 0l1.232-4l1.24 3.998L12.815 25l4-.003l-3.234 2.474z"/><circle cx="24.75" cy="30" r=".5" fill="#fff"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}