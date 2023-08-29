package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightSafety(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFlightSafety0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M38.986 23c-.313 10.709-5.866 16.522-13.984 20c-4.393-1.881-8.034-4.447-10.502-8.101M39 19v-8.744L25.007 5L11 10.256v11.773c0 2.587.294 4.9.848 6.971"/><path fill="#fff" stroke-linecap="round" d="M9.268 30.632c-1.373-1.257-2.99-3.877-3.587-4.027c0 0-1.798 4.434-1.68 7.395c.117 2.96 2.896 4.635 5.824 3.22C12.753 35.806 44 20 44 20l-9-2L9.268 30.632Z"/><path stroke-linecap="round" d="m28 21l-9-3l-2 1l3 6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFlightSafety0)"/>`),
		g.Group(children),
	)
}