package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipHorizontalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m3.5.5l.224-.447A.5.5 0 0 0 3 .5h.5Zm8 4V5a.5.5 0 0 0 .224-.947L11.5 4.5Zm-8 0H3a.5.5 0 0 0 .5.5v-.5Zm0 6V10a.5.5 0 0 0-.5.5h.5Zm8 0l.224.447A.5.5 0 0 0 11.5 10v.5Zm-8 4H3a.5.5 0 0 0 .724.447L3.5 14.5ZM3.276.947l8 4l.448-.894l-8-4l-.448.894ZM11.5 4h-8v1h8V4ZM4 4.5v-4H3v4h1ZM0 8h15V7H0v1Zm3.5 3h8v-1h-8v1Zm7.776-.947l-8 4l.448.894l8-4l-.448-.894ZM4 14.5v-4H3v4h1Z"/>`),
		g.Group(children),
	)
}