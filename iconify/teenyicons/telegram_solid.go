package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelegramSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14.993 1.582a.5.5 0 0 0-.661-.553l-14 5a.5.5 0 0 0-.056.918l4 2a.5.5 0 0 0 .501-.031l3.32-2.214L6.11 9.188a.5.5 0 0 0 .113.728l6 4a.5.5 0 0 0 .77-.334l2-12Z"/>`),
		g.Group(children),
	)
}