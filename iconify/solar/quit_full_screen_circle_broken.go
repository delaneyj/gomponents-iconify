package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuitFullScreenCircleBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M9.998 6c-.015 1.707-.11 2.648-.73 3.268c-.62.62-1.561.716-3.268.73m3.998 7.999c-.015-1.706-.11-2.647-.73-3.267c-.62-.62-1.561-.716-3.268-.73m8-8c.014 1.707.11 2.648.73 3.268c.62.62 1.56.716 3.267.73M14 17.997c.014-1.706.11-2.647.73-3.267c.62-.62 1.56-.716 3.267-.73"/><path d="M7 3.338A9.954 9.954 0 0 1 12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12c0-1.821.487-3.53 1.338-5"/></g>`),
		g.Group(children),
	)
}