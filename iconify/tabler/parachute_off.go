package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParachuteOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 12c0-5.523-4.477-10-10-10c-1.737 0-3.37.443-4.794 1.222m-2.28 1.71A9.969 9.969 0 0 0 2 12"/><path d="M22 12c0-1.66-1.46-3-3.25-3c-1.63 0-2.973 1.099-3.212 2.54m-.097-.09c-.23-1.067-1.12-1.935-2.29-2.284m-3.445.568C8.967 10.284 8.5 11.094 8.5 12c0-1.66-1.46-3-3.25-3C3.45 9 2 10.34 2 12m0 0l10 10l-3.5-10m6.082 2.624L12 22l4.992-4.992m2.014-2.014l3-3M3 3l18 18"/></g>`),
		g.Group(children),
	)
}