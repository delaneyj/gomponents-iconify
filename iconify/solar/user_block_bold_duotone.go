package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserBlockBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 10a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path fill-rule="evenodd" d="M16.5 15.75a2.75 2.75 0 0 0-2.383 4.123l3.756-3.756a2.735 2.735 0 0 0-1.373-.367Zm2.42 1.442l-3.728 3.728a2.75 2.75 0 0 0 3.728-3.728ZM12.25 18.5a4.25 4.25 0 1 1 8.5 0a4.25 4.25 0 0 1-8.5 0Z" clip-rule="evenodd"/><path d="M17.996 14.521a4.25 4.25 0 0 0-3.979 7.429c-.608.033-1.278.05-2.017.05c-8 0-8-2.015-8-4.5S7.582 13 12 13c2.387 0 4.53.588 5.996 1.521Z" opacity=".4"/></g>`),
		g.Group(children),
	)
}