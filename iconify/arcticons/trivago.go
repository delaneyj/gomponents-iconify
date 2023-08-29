package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trivago(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.313 21.575l-2 5.3l-2-5.3m15.69 0v6a2.006 2.006 0 0 1-2 2h0a1.678 1.678 0 0 1-1.4-.6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.003 21.575h0a2.006 2.006 0 0 1 2 2v1.3a2.006 2.006 0 0 1-2 2h0a2.006 2.006 0 0 1-2-2v-1.3a2.006 2.006 0 0 1 2-2Zm5.998 5.3h0a2.006 2.006 0 0 1-2-2v-1.3a2.006 2.006 0 0 1 2-2h0a2.006 2.006 0 0 1 2 2v1.3a2.006 2.006 0 0 1-2 2Zm-25.705-3.299a2.006 2.006 0 0 1 2-2h0m-2 0v5.299"/><circle cx="16.306" cy="19.175" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.306 21.575v5.3m-7.307-7v6a.945.945 0 0 0 1 1h.3m-2.3-5.3h2.1m17.9 3.3a2.006 2.006 0 0 1-2 2h0a2.006 2.006 0 0 1-2-2v-1.3a2.006 2.006 0 0 1 2-2h0a2.006 2.006 0 0 1 2 2m0 3.3v-5.3"/>`),
		g.Group(children),
	)
}