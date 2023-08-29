package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 3V1.75c-.69 0-1.25.56-1.25 1.25H3Zm15.293 11.293l-1.139-.515a1.25 1.25 0 0 0 .255 1.399l.884-.884ZM21 17l.884.884a1.25 1.25 0 0 0 0-1.768L21 17Zm-4 4l-.884.884a1.25 1.25 0 0 0 1.768 0L17 21Zm-2.707-2.707l.884-.884a1.25 1.25 0 0 0-1.4-.255l.516 1.139ZM11 1.75H3v2.5h8v-2.5ZM20.25 11A9.25 9.25 0 0 0 11 1.75v2.5A6.75 6.75 0 0 1 17.75 11h2.5Zm-.818 3.808A9.22 9.22 0 0 0 20.25 11h-2.5a6.72 6.72 0 0 1-.596 2.778l2.278 1.03Zm-2.023.369l2.707 2.707l1.768-1.768l-2.707-2.707l-1.768 1.768Zm2.707.94l-4 4l1.768 1.767l4-4l-1.768-1.768Zm-2.232 4l-2.707-2.708l-1.768 1.768l2.707 2.707l1.768-1.768ZM11 20.25a9.222 9.222 0 0 0 3.808-.818l-1.03-2.278A6.721 6.721 0 0 1 11 17.75v2.5ZM1.75 11A9.25 9.25 0 0 0 11 20.25v-2.5A6.75 6.75 0 0 1 4.25 11h-2.5Zm0-8v8h2.5V3h-2.5Z"/><circle cx="11" cy="11" r="2" fill="currentColor" transform="rotate(-180 11 11)"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m3 3l8 8"/></g>`),
		g.Group(children),
	)
}