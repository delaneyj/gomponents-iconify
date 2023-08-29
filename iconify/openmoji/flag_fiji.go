package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagFiji(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#61b2e4" d="M5 17h62v38H5z"/><path fill="#1e50a0" d="M5 17h31v19H5z"/><path fill="#d22f27" d="M52 48s6-1.938 6-7.75V34.5H46v5.75c0 5.813 6 7.75 6 7.75Z"/><path fill="#fff" d="M53 37h5v4h-5zm-7 0h5v4h-5zm5 10.574V43h-4.488A9.028 9.028 0 0 0 51 47.574ZM57.488 43H53v4.574A9.028 9.028 0 0 0 57.488 43Z"/><path fill="#d22f27" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" d="M46.5 34h11v2h-11z"/><path fill="#fff" d="M9.887 18H6v2.332L32.113 36H36v-2.332L9.887 18z"/><path fill="#fff" d="M36 20.332V18h-3.887L6 33.668V36h3.887L36 20.332z"/><path fill="#fff" d="M6 24h30v6H6z"/><path fill="#fff" d="M18 18h6v18h-6z"/><path fill="#d22f27" d="M20 18h2v18h-2z"/><path fill="#d22f27" d="M6 26h30v2H6zm30 7.668L29.887 30H26l10 6v-2.332zM36 18h-3.887L24 22.868V24h2l10-6zM6 20.332L12.113 24H16L6 18v2.332zM6 36h3.887L18 31.132V30h-2L6 36z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}