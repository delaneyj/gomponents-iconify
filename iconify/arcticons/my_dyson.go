package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyDyson(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.016 24.65v2.7a2 2 0 0 1-2 2h0a1.993 1.993 0 0 1-1.414-.586"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.016 21.35v3.3a2 2 0 0 1-2 2h0a2 2 0 0 1-2-2v-3.3"/><rect width="4" height="5.3" x="28.048" y="21.35" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.48 26.203c.365.306.76.447 1.645.447h.449c.73 0 1.322-.593 1.322-1.325h0c0-.732-.592-1.325-1.322-1.325h-.898c-.73 0-1.322-.593-1.322-1.325h0c0-.732.592-1.325 1.322-1.325h.45c.885 0 1.279.14 1.644.447m12.57 4.853v-3.3a2 2 0 0 0-2-2h0a2 2 0 0 0-2 2m0 3.3v-5.3m-20.68 2a2 2 0 0 0-2-2h0a2 2 0 0 0-2 2v1.3a2 2 0 0 0 2 2h0a2 2 0 0 0 2-2m0 2v-8"/>`),
		g.Group(children),
	)
}