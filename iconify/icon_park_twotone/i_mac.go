package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IMac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTIMac0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M4 6h40v22H4V6Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 28v8H4v-8"/><path fill="#fff" d="M13.09 18h-2.18a.91.91 0 0 0-.91.91v2.18c0 .503.407.91.91.91h2.18a.91.91 0 0 0 .91-.91v-2.18a.91.91 0 0 0-.91-.91Zm0-6h-2.18a.91.91 0 0 0-.91.91v2.18c0 .503.407.91.91.91h2.18a.91.91 0 0 0 .91-.91v-2.18a.91.91 0 0 0-.91-.91Zm8 6h-2.18a.91.91 0 0 0-.91.91v2.18c0 .503.407.91.91.91h2.18a.91.91 0 0 0 .91-.91v-2.18a.91.91 0 0 0-.91-.91Zm0-6h-2.18a.91.91 0 0 0-.91.91v2.18c0 .503.407.91.91.91h2.18a.91.91 0 0 0 .91-.91v-2.18a.91.91 0 0 0-.91-.91Zm8 6h-2.18a.91.91 0 0 0-.91.91v2.18c0 .503.407.91.91.91h2.18a.91.91 0 0 0 .91-.91v-2.18a.91.91 0 0 0-.91-.91Zm0-6h-2.18a.91.91 0 0 0-.91.91v2.18c0 .503.407.91.91.91h2.18a.91.91 0 0 0 .91-.91v-2.18a.91.91 0 0 0-.91-.91Zm8 0h-2.18a.91.91 0 0 0-.91.91v2.18c0 .503.407.91.91.91h2.18a.91.91 0 0 0 .91-.91v-2.18a.91.91 0 0 0-.91-.91Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20.846 36L16 42h16l-4.846-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTIMac0)"/>`),
		g.Group(children),
	)
}