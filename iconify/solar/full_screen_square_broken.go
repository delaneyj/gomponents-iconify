package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullScreenSquareBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M6 9.997c.014-1.706.11-2.647.73-3.267c.62-.62 1.56-.716 3.267-.73M6 14c.014 1.707.11 2.648.73 3.268c.62.62 1.56.716 3.267.73m8.001-8.001c-.015-1.706-.11-2.647-.73-3.267c-.62-.62-1.561-.716-3.268-.73m3.998 8c-.015 1.707-.11 2.648-.73 3.268c-.62.62-1.561.716-3.268.73"/><path d="M22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12s0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464c.974.974 1.3 2.343 1.41 4.536"/></g>`),
		g.Group(children),
	)
}