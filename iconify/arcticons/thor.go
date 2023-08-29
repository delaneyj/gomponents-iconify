package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 22.19v-1.383h6.387v1.383m-3.194-1.383v6.067m-1.277.319h2.555m4.258-6.386h1.383v6.067H16.23M17.612 24h4.684m1.384-3.193h-1.384v6.067h1.384"/><ellipse cx="28.736" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.3" ry="2.981"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.024 27.193v-6.067m-1.278 0h4.655a1.49 1.49 0 1 1 0 2.98h-3.377m2.815 0l1.597 3.087H40.5m-6.54 0h2.554"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}