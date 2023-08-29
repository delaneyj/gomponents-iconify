package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailOpenLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.243 6.854l9.247-5.546a1 1 0 0 1 1.028 0l9.24 5.546a.5.5 0 0 1 .242.429V20a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V7.283a.5.5 0 0 1 .243-.429ZM4 8.132V19h16V8.132l-7.996-4.8L4 8.132Zm8.06 5.566l5.296-4.463l1.288 1.53l-6.57 5.537l-6.71-5.53l1.272-1.544l5.424 4.47Z"/>`),
		g.Group(children),
	)
}