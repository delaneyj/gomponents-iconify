package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmazonPrimeVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.993 20.125l-2 5.3l-2-5.3"/><rect width="4" height="5.3" x="32.007" y="20.215" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><circle cx="18.007" cy="17.675" r=".7" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.007 20.125v5.3m11.738-1.009a2 2 0 0 1-1.738 1.009h0a2 2 0 0 1-2-2v-1.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2v.65h-4m-2-.65a2 2 0 0 0-2-2h0a2 2 0 0 0-2 2v1.3a2 2 0 0 0 2 2h0a2 2 0 0 0 2-2m0 2v-8"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.28 29.406c1.113-.45 3.092-1.05 3.688-.327c.644.781-.17 2.477-.92 3.794"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.798 29.928c1.759 1.397 6.954 3.535 12.488 3.535a17.003 17.003 0 0 0 10.167-3.08"/>`),
		g.Group(children),
	)
}