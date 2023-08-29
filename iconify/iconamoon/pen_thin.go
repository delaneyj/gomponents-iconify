package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 3v-.5a.5.5 0 0 0-.5.5H3Zm15.293 11.293l-.456-.206a.5.5 0 0 0 .102.56l.354-.354ZM21 17l.354.354a.5.5 0 0 0 0-.708L21 17Zm-4 4l-.354.354a.5.5 0 0 0 .708 0L17 21Zm-2.707-2.707l.354-.354a.5.5 0 0 0-.56-.102l.206.456ZM11 2.5H3v1h8v-1Zm8.5 8.5A8.5 8.5 0 0 0 11 2.5v1a7.5 7.5 0 0 1 7.5 7.5h1Zm-.751 3.499A8.472 8.472 0 0 0 19.5 11h-1a7.471 7.471 0 0 1-.663 3.087l.912.412Zm-.81.147l2.707 2.708l.708-.708l-2.707-2.707l-.708.707Zm2.707 2l-4 4l.708.708l4-4l-.708-.708Zm-3.292 4l-2.707-2.707l-.708.707l2.707 2.708l.708-.708ZM11 19.5a8.472 8.472 0 0 0 3.499-.751l-.412-.912A7.471 7.471 0 0 1 11 18.5v1ZM2.5 11a8.5 8.5 0 0 0 8.5 8.5v-1A7.5 7.5 0 0 1 3.5 11h-1Zm0-8v8h1V3h-1Z"/><circle cx="11" cy="11" r="1" stroke="currentColor" transform="rotate(-180 11 11)"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m3 3l7 7"/></g>`),
		g.Group(children),
	)
}