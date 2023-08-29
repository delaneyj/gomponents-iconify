package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FluxTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="17.343" cy="24.701" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="12.843" ry="12.87"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.859 14.068a7.839 7.839 0 0 1 13.449 8.044"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.813 21.613h5.712a7.976 7.976 0 0 1 0 15.95H16.354"/>`),
		g.Group(children),
	)
}