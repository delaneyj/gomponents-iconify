package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagBrazil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5c9e31" d="M5 17h62v38H5z"/><path fill="#fcea2b" d="m59.023 36.023l-23.157 14.63l-22.889-14.362l23.157-14.63l22.889 14.362z"/><circle cx="36" cy="36" r="9" fill="#1e50a0"/><path fill="#fff" d="M44.159 39.782a9.046 9.046 0 0 0 .696-2.26a11.473 11.473 0 0 0-17.477-4.04a8.984 8.984 0 0 0-.352 2.013a10.998 10.998 0 0 1 17.133 4.287Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}