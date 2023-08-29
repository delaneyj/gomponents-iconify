package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopCalendarCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" transform="translate(3 3)"><path fill-rule="evenodd" d="M3 4h14a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1Zm1 4v8h12V8H4Z" clip-rule="evenodd"/><circle cx="6.5" cy="10.5" r="1.5"/><circle cx="5.5" cy="4.5" r="1.5"/><circle cx="14.5" cy="4.5" r="1.5"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopCalendarCircleFilled0)"/></g>`),
		g.Group(children),
	)
}