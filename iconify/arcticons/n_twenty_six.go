package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NTwentySix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.79 18.85v10.46m7.02 0V18.85m-7.02 0l7.02 10.46m17.7-10.48a3.65 3.65 0 0 0-3.32 3.71v3.26a3.51 3.51 0 0 0 3.51 3.51h0a3.51 3.51 0 0 0 3.51-3.51"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.21 25.8a3.52 3.52 0 0 0-3.51-3.52h0a3.52 3.52 0 0 0-3.51 3.52m3.32-6.97h1.85m-8.89 3.51a3.5 3.5 0 0 0-3.7-3.51a3.64 3.64 0 0 0-3.32 3.7m7.02 6.78h-7l5.29-3.67a3.81 3.81 0 0 0 1.6-2.12a3.64 3.64 0 0 0 .13-1m-17.7 9.24h7.02M9.79 16.24h7.02"/><rect width="38" height="38" x="5" y="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`),
		g.Group(children),
	)
}