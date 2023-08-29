package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FootballFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2Zm1.67 14h-3.34l-1.38 1.897l.554 1.706A7.993 7.993 0 0 0 12 20c.871 0 1.71-.14 2.496-.397l.553-1.706L13.669 16Zm-8.376-5.128l-1.292.938L4 12c0 1.73.549 3.331 1.482 4.64h1.91l1.323-1.82l-1.028-3.17l-2.393-.778Zm13.412 0l-2.393.778l-1.028 3.17l1.322 1.82h1.91A7.963 7.963 0 0 0 20 12l-.003-.191l-1.291-.937ZM14.29 4.333l-1.29.94V7.79l2.694 1.957l2.24-.727l.554-1.703a8.014 8.014 0 0 0-4.196-2.984Zm-4.582 0a8.014 8.014 0 0 0-4.196 2.985l.554 1.702l2.239.727L11 7.79V5.273l-1.291-.94Z"/>`),
		g.Group(children),
	)
}