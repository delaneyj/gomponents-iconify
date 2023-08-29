package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailRewindTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.998 7.999V6.8l2.36 1.966a1 1 0 0 0 1.64-.768V3.002a1 1 0 0 0-1.64-.768L15.998 4.2V3.002a1 1 0 0 0-1.64-.768L11.36 4.732a1 1 0 0 0 0 1.537l2.998 2.498a1 1 0 0 0 1.64-.768Zm.838.801c-.458 1.058-1.792 1.55-2.842.93l-3.74 2.201a.5.5 0 0 1-.507 0L2 7.373V14.5A2.5 2.5 0 0 0 4.5 17h11a2.5 2.5 0 0 0 2.5-2.5V9.734a2.085 2.085 0 0 1-.283-.199l-.881-.735ZM4.5 4h6.177a2 2 0 0 0 .043 3.037l2.433 2.028L10 10.92L2.015 6.223A2.5 2.5 0 0 1 4.5 4Z"/>`),
		g.Group(children),
	)
}