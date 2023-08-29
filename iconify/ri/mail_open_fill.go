package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailOpenFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.243 6.854l9.247-5.546a1 1 0 0 1 1.028 0l9.24 5.546a.5.5 0 0 1 .242.429V20a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V7.283a.5.5 0 0 1 .243-.429Zm16.103 1.39l-6.285 5.439l-6.414-5.445l-1.294 1.524l7.72 6.555l7.581-6.561l-1.308-1.512Z"/>`),
		g.Group(children),
	)
}