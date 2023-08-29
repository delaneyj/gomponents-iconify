package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowChartSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M9.8 18.8h16.4v3.08h1.6V17.2h-9V14h-1.6v3.2h-9v4.68h1.6V18.8z" class="clr-i-solid clr-i-solid-path-1"/><rect width="14" height="10" x="2" y="23" fill="currentColor" class="clr-i-solid clr-i-solid-path-2" rx="2" ry="2"/><rect width="14" height="10" x="20" y="23" fill="currentColor" class="clr-i-solid clr-i-solid-path-3" rx="2" ry="2"/><rect width="14" height="10" x="11" y="3" fill="currentColor" class="clr-i-solid clr-i-solid-path-4" rx="2" ry="2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}