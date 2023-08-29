package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoFramePlayHorizontalLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 2c4.714 0 7.071 0 8.535 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12s0-7.071 1.464-8.536C4.93 2 7.286 2 12 2Z"/><path stroke-linecap="round" d="M21.5 17h-19m19-10h-19M12 17v5m0-20v5m5-4.5V7m0 10v4.5M7 17v4.5m0-19V7" opacity=".5"/><path stroke-linecap="round" d="M14 12c0-.528-.53-.884-1.589-1.596c-1.073-.721-1.61-1.082-2.01-.817C10 9.852 10 10.567 10 12c0 1.432 0 2.148.4 2.413c.4.265.938-.096 2.011-.817C13.471 12.884 14 12.528 14 12Z"/></g>`),
		g.Group(children),
	)
}