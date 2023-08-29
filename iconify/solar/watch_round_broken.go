package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchRoundBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m17 6.5l-.304-1.368c-.334-1.501-.5-2.252-1.049-2.692C15.1 2 14.33 2 12.791 2H11.21c-1.54 0-2.31 0-2.857.44c-.549.44-.715 1.19-1.05 2.692L7 6.5m10 11l-.304 1.368c-.334 1.501-.5 2.252-1.049 2.692c-.548.44-1.317.44-2.856.44H11.21c-1.539 0-2.308 0-2.856-.44c-.549-.44-.715-1.19-1.05-2.692L7 17.5"/><path stroke-linecap="round" d="M4.755 10.058a7.5 7.5 0 1 1 0 3.884"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 8.923V12l2 1.923"/></g>`),
		g.Group(children),
	)
}