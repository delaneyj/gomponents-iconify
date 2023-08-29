package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M3 3h8a8 8 0 0 1 7.293 11.293L21 17l-4 4l-2.707-2.707A8 8 0 0 1 3 11V3Z" clip-rule="evenodd" opacity=".16"/><path fill="currentColor" d="M3 3V2a1 1 0 0 0-1 1h1Zm15.293 11.293l-.911-.412a1 1 0 0 0 .204 1.12l.707-.708ZM21 17l.707.707a1 1 0 0 0 0-1.414L21 17Zm-4 4l-.707.707a1 1 0 0 0 1.414 0L17 21Zm-2.707-2.707l.707-.707a1 1 0 0 0-1.12-.204l.413.911ZM11 2H3v2h8V2Zm9 9a9 9 0 0 0-9-9v2a7 7 0 0 1 7 7h2Zm-.796 3.705A8.971 8.971 0 0 0 20 11h-2a6.97 6.97 0 0 1-.618 2.88l1.822.825ZM17.586 15l2.707 2.707l1.414-1.414L19 13.586L17.586 15Zm2.707 1.293l-4 4l1.414 1.414l4-4l-1.414-1.414Zm-2.586 4L15 17.586L13.586 19l2.707 2.707l1.414-1.414ZM11 20a8.971 8.971 0 0 0 3.705-.796l-.824-1.822A6.97 6.97 0 0 1 11 18v2Zm-9-9a9 9 0 0 0 9 9v-2a7 7 0 0 1-7-7H2Zm0-8v8h2V3H2Z"/><circle cx="11" cy="11" r="2" fill="currentColor" transform="rotate(-180 11 11)"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l8 8"/></g>`),
		g.Group(children),
	)
}