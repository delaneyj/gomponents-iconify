package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarAssistantTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 5.5A2.5 2.5 0 0 0 14.5 3h-9A2.5 2.5 0 0 0 3 5.5v9A2.5 2.5 0 0 0 5.5 17h4.1a5.465 5.465 0 0 1-.393-1H5.5A1.5 1.5 0 0 1 4 14.5V7h12v2.207c.349.099.683.23 1 .393V5.5ZM5.5 4h9A1.5 1.5 0 0 1 16 5.5V6H4v-.5A1.5 1.5 0 0 1 5.5 4ZM19 14.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Zm-5.12-3a.5.5 0 0 0-.472.334l-.88 2.5A.5.5 0 0 0 13 15h.776l-.508 1.869a.5.5 0 0 0 .903.401l2.25-3.5A.5.5 0 0 0 16 13h-.746l.28-.842a.5.5 0 0 0-.474-.658h-1.18Z"/>`),
		g.Group(children),
	)
}