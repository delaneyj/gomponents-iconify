package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquintingFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm5.106-3.947a1 1 0 0 0 .447 1.341l.21.106l-.21.106a1 1 0 1 0 .894 1.788l2-1a1 1 0 0 0 0-1.788l-2-1a1 1 0 0 0-1.341.447Zm9.341 1.341a1 1 0 1 0-.894-1.788l-2 1a1 1 0 0 0 0 1.788l2 1a1 1 0 1 0 .894-1.788l-.211-.106l.211-.106ZM8.535 13a1 1 0 0 0-.865 1.5A4.998 4.998 0 0 0 12 17a4.998 4.998 0 0 0 4.33-2.5a1 1 0 0 0-.865-1.5h-6.93Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}