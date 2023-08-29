package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilCalendarCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M19.5 7.5h-13A.5.5 0 0 0 6 8v11a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-.5-.5ZM7 18.5v-10h12v10H7Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M19.5 7.5h-13A.5.5 0 0 0 6 8v3a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-.5-.5ZM7 10.5v-2h12v2H7Z" clip-rule="evenodd"/><path d="M8.5 8.5h1A.5.5 0 0 0 10 8V7a.5.5 0 0 0-.5-.5h-1A.5.5 0 0 0 8 7v1a.5.5 0 0 0 .5.5Zm8 0h1A.5.5 0 0 0 18 8V7a.5.5 0 0 0-.5-.5h-1a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5Zm-7.5 6a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilCalendarCircleFilled0)"/></g>`),
		g.Group(children),
	)
}