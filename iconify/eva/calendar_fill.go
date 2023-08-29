package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaCalendarFill0"><g id="evaCalendarFill1"><path id="evaCalendarFill2" fill="currentColor" d="M18 4h-1V3a1 1 0 0 0-2 0v1H9V3a1 1 0 0 0-2 0v1H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3ZM8 17a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm8 0h-4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2Zm3-6H5V7a1 1 0 0 1 1-1h1v1a1 1 0 0 0 2 0V6h6v1a1 1 0 0 0 2 0V6h1a1 1 0 0 1 1 1Z"/></g></g>`),
		g.Group(children),
	)
}