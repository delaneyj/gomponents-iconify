package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PreviousTrack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M16.5 14.796L11.019 10L16.5 5.204v9.592Zm-6.14-4.043a1 1 0 0 1 0-1.506l5.482-4.796c.646-.566 1.658-.106 1.658.753v9.592a1 1 0 0 1-1.659.753l-5.48-4.796Z"/><path d="M9.5 14.796L4.019 10L9.5 5.204v9.592Zm-6.14-4.043a1 1 0 0 1 0-1.506l5.482-4.796c.646-.566 1.658-.106 1.658.753v9.592c0 .86-1.012 1.319-1.658.753L3.36 10.753Z"/><path d="M3 4.5a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 1 0V5a.5.5 0 0 0-.5-.5Z"/></g>`),
		g.Group(children),
	)
}