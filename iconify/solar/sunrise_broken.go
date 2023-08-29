package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunriseBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 22h8M5 19h14M2 16h20"/><path fill="currentColor" d="m12 10l.53-.53a.75.75 0 0 0-1.06 0L12 10Zm1.47 2.53a.75.75 0 1 0 1.06-1.06l-1.06 1.06Zm-4-1.06a.75.75 0 1 0 1.06 1.06l-1.06-1.06ZM12.75 16v-6h-1.5v6h1.5Zm-1.28-5.47l2 2l1.06-1.06l-2-2l-1.06 1.06Zm0-1.06l-2 2l1.06 1.06l2-2l-1.06-1.06Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 2v1m10 9h-1M3 12H2m17.07-7.07l-.392.393M5.322 5.322l-.393-.393M10 6.341a6 6 0 0 1 6.5 9.627h-9A5.977 5.977 0 0 1 6 12c0-.701.12-1.374.341-2"/></g>`),
		g.Group(children),
	)
}