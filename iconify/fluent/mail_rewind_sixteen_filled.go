package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailRewindSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.998 6.999V5.8l2.36 1.966a1 1 0 0 0 1.64-.768V2.002a1 1 0 0 0-1.64-.768L11.998 3.2V2.002a1 1 0 0 0-1.64-.768L7.36 3.732a1 1 0 0 0 0 1.537l2.998 2.498a1 1 0 0 0 1.64-.768Zm.722 1.028l.235-.127l.762.635c.092.076.186.142.283.2V12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7.337l5.763 3.103a.5.5 0 0 0 .474 0l2.673-1.439a2.032 2.032 0 0 0 1.81-.974ZM4 4h2.063a1.998 1.998 0 0 0 .657 2.037l2.977 2.481L8 9.432l-6-3.23V6a2 2 0 0 1 2-2Z"/>`),
		g.Group(children),
	)
}