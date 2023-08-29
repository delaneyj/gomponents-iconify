package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClusterSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31.36 8H27.5v24H33V9.67A1.65 1.65 0 0 0 31.36 8Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M3 9.67V32h5.5V8H4.64A1.65 1.65 0 0 0 3 9.67Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="currentColor" d="M24.32 4H11.68A1.68 1.68 0 0 0 10 5.68V32h16V5.68A1.68 1.68 0 0 0 24.32 4ZM18 27.79A1.79 1.79 0 1 1 19.81 26A1.8 1.8 0 0 1 18 27.79Zm5-17.19H13V9h10Z" class="clr-i-solid clr-i-solid-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}