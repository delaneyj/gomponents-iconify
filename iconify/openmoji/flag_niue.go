package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagNiue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fcea2b" d="M5 17h62v38H5z"/><path fill="#1e50a0" d="M5 17h31v19H5z"/><path fill="#fff" d="M9.887 18H6v2.332L32.113 36H36v-2.332L9.887 18z"/><path fill="#fff" d="M36 20.332V18h-3.887L6 33.668V36h3.887L36 20.332z"/><path fill="#fff" d="M6 24h30v6H6z"/><path fill="#fff" d="M18 18h6v18h-6z"/><path fill="#d22f27" d="M20 18h2v18h-2z"/><path fill="#d22f27" d="M6 26h30v2H6zm30 7.668L29.887 30H26l10 6v-2.332zM36 18h-3.887L24 22.868V24h2l10-6zM6 20.332L12.113 24H16L6 18v2.332zM6 36h3.887L18 31.132V30h-2L6 36z"/><circle cx="20.5" cy="26.5" r="2.5" fill="#1e50a0"/><path fill="#fcea2b" stroke="#fcea2b" stroke-linecap="round" stroke-linejoin="round" d="m19.531 28l.995-3l.859 2.954L19 26.174l3-.074l-2.469 1.9z"/><circle cx="27" cy="26.5" r="1" fill="#fcea2b"/><circle cx="14" cy="26.5" r="1" fill="#fcea2b"/><circle cx="20.5" cy="32" r="1" fill="#fcea2b"/><circle cx="20.5" cy="21" r="1" fill="#fcea2b"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}