package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuitFullScreenCircleLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M9.998 6c-.015 1.707-.11 2.648-.73 3.268c-.62.62-1.561.716-3.268.73m3.998 7.999c-.015-1.706-.11-2.647-.73-3.267c-.62-.62-1.561-.716-3.268-.73m8-8c.014 1.707.11 2.648.73 3.268c.62.62 1.56.716 3.267.73M14 17.997c.014-1.706.11-2.647.73-3.267c.62-.62 1.56-.716 3.267-.73"/><circle cx="12" cy="12" r="10"/></g>`),
		g.Group(children),
	)
}