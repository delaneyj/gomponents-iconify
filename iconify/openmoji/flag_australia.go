package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagAustralia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#1e50a0" d="M5 17h62v38H5z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m54.233 38.945l.927-3l.927 3l-2.427-1.855l3 .001l-2.427 1.854zM20.5 46.999l-1.558 1.477l.155-2.17L17 45.978l1.75-1.23l-1.057-1.886l2.028.637l.779-2.023l.779 2.023l2.028-.637l-1.057 1.886l1.75 1.23l-2.097.328l.155 2.17l-1.558-1.477zm24-11.578l-1.113 1.055l.111-1.55L42 34.691l1.25-.878l-.755-1.347l1.449.455l.556-1.445l.556 1.445l1.449-.455l-.755 1.347l1.25.878l-1.498.235l.111 1.55l-1.113-1.055zm15-3l-1.113 1.055l.111-1.55L57 31.691l1.25-.878l-.755-1.347l1.449.455l.556-1.445l.556 1.445l1.449-.455l-.755 1.347l1.25.878l-1.498.235l.111 1.55l-1.113-1.055zm-8-6l-1.113 1.055l.111-1.55L49 25.691l1.25-.878l-.755-1.347l1.449.455l.556-1.445l.556 1.445l1.449-.455l-.755 1.347l1.25.878l-1.498.235l.111 1.55l-1.113-1.055zm0 21l-1.113 1.055l.111-1.55L49 46.691l1.25-.878l-.755-1.347l1.449.455l.556-1.445l.556 1.445l1.449-.455l-.755 1.347l1.25.878l-1.498.235l.111 1.55l-1.113-1.055z"/><path fill="#fff" d="M9.887 18H6v2.332L32.113 36H36v-2.332L9.887 18z"/><path fill="#fff" d="M36 20.332V18h-3.887L6 33.668V36h3.887L36 20.332z"/><path fill="#fff" d="M6 24h30v6H6z"/><path fill="#fff" d="M18 18h6v18h-6z"/><path fill="#d22f27" d="M20 18h2v18h-2z"/><path fill="#d22f27" d="M6 26h30v2H6zm30 7.668L29.887 30H26l10 6v-2.332zM36 18h-3.887L24 22.868V24h2l10-6zM6 20.332L12.113 24H16L6 18v2.332zM6 36h3.887L18 31.132V30h-2L6 36z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}