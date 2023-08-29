package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailRewindSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.998 6.999V5.8l2.36 1.966a1 1 0 0 0 1.64-.768V2.002a1 1 0 0 0-1.64-.768L11.998 3.2V2.002a1 1 0 0 0-1.64-.768L7.36 3.732a1 1 0 0 0 0 1.537l2.998 2.498a1 1 0 0 0 1.64-.768Zm-3.76 3.441l2.672-1.439a1.97 1.97 0 0 1-1.192-.466l-.02-.017L8 9.432L3 6.74V6a1 1 0 0 1 1-1h2.063a2.016 2.016 0 0 1 0-1H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V8.734a2.092 2.092 0 0 1-.283-.199L13 7.937V12a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V7.876l4.763 2.564a.5.5 0 0 0 .474 0Z"/>`),
		g.Group(children),
	)
}