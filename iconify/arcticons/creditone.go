package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Creditone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.88 32.72a2.4 2.4 0 0 1-1.87 1h0a1.58 1.58 0 0 1-1.6-2l.22-1.26a2.42 2.42 0 0 1 2.29-1.95h0a1.58 1.58 0 0 1 1.6 1.95l-.11.63h-3.89m-2.45 2.61l.57-3.21A1.59 1.59 0 0 0 23 28.54h0a2.43 2.43 0 0 0-2.29 1.95l-.57 3.21m.6-3.21l.34-1.95m-7.6 2.58a2.1 2.1 0 0 0 2.12 2.58a3.21 3.21 0 0 0 3-2.58l.46-2.63A2.11 2.11 0 0 0 17 25.91a3.22 3.22 0 0 0-3 2.58Zm4.46-11.81a2.43 2.43 0 0 1 2.3-1.94h0m-1.95 0l-.91 5.16"/><ellipse cx="33.85" cy="14.98" fill="currentColor" rx=".74" ry=".62" transform="rotate(-42.48 33.848 14.977)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.43 17.37l-.91 5.16m-7.98-.98a2.41 2.41 0 0 1-1.87 1h0a1.59 1.59 0 0 1-1.6-2l.22-1.27a2.43 2.43 0 0 1 2.29-1.94h0a1.59 1.59 0 0 1 1.61 1.94l-.11.64h-3.9m15.49-4.16l-1 5.79a.81.81 0 0 0 .81 1h.29m-1.41-5.18h2.04m-21.6 2.55v0a3.22 3.22 0 0 1-3 2.58h0a2.1 2.1 0 0 1-2.16-2.5l.46-2.63a3.21 3.21 0 0 1 3-2.58h0a2.1 2.1 0 0 1 2.12 2.58v0m14.83 1.94a1.59 1.59 0 0 0-1.61-1.94h0a2.42 2.42 0 0 0-2.29 1.94l-.22 1.27a1.58 1.58 0 0 0 1.6 2h0a2.42 2.42 0 0 0 2.29-2m-.34 1.95l1.37-7.79"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}