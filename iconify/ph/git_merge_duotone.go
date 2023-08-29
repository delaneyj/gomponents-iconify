package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitMergeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M104 56a24 24 0 1 1-24-24a24 24 0 0 1 24 24Z" opacity=".2"/><path d="M208 112a32.06 32.06 0 0 0-31 24h-25a40.19 40.19 0 0 1-32-16L93.69 84.92A32 32 0 1 0 72 87v82a32 32 0 1 0 16 0v-65l19.2 25.6A56.26 56.26 0 0 0 152 152h25a32 32 0 1 0 31-40ZM64 56a16 16 0 1 1 16 16a16 16 0 0 1-16-16Zm32 144a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm112-40a16 16 0 1 1 16-16a16 16 0 0 1-16 16Z"/></g>`),
		g.Group(children),
	)
}