package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScatterPlotSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M34 12.34V29a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h18.57a7.44 7.44 0 0 0 6.74 8.46l1.01.02A7.45 7.45 0 0 0 34 12.34ZM9.101 15.8a.8.8 0 0 0 1.13 0l1.16-1.16l1.16 1.16a.8.8 0 1 0 1.13-1.13l-1.15-1.13l1.16-1.16a.8.8 0 1 0-1.13-1.13l-1.16 1.16l-1.17-1.19a.8.8 0 1 0-1.13 1.13l1.16 1.19l-1.16 1.13a.8.8 0 0 0 0 1.13Zm6.075 9.736a.8.8 0 0 0 1.13 0l1.16-1.16l1.16 1.16a.8.8 0 1 0 1.13-1.13l-1.15-1.13l1.16-1.16a.8.8 0 1 0-1.13-1.13l-1.16 1.16l-1.17-1.19a.8.8 0 1 0-1.13 1.13l1.16 1.19l-1.16 1.13a.8.8 0 0 0 0 1.13Zm7.736-5.193a.8.8 0 0 0 1.13 0l1.16-1.16l1.16 1.16a.8.8 0 1 0 1.13-1.13l-1.15-1.13l1.16-1.16a.8.8 0 1 0-1.13-1.13l-1.16 1.16l-1.17-1.19a.8.8 0 1 0-1.13 1.13l1.16 1.19l-1.16 1.13a.8.8 0 0 0 0 1.13Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-2--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}