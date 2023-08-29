package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m24 14.358l-3.04-2.965l2.608-3.348l-4.114-1.051l.584-4.204l-4.088 1.142L14.35 0l-2.965 3.04L8.037.432L6.986 4.546l-4.204-.584l1.142 4.087l-3.932 1.596l3.04 2.966l-2.608 3.348l4.114 1.051l-.59 4.204l4.087-1.142l1.6 3.932l2.965-3.04l3.348 2.608l1.051-4.114l4.205.59l-1.142-4.087zm-9.719 4.302a7.04 7.04 0 1 1 4.378-8.99l.015.049c.24.679.378 1.461.378 2.276a7.042 7.042 0 0 1-4.722 6.649l-.049.015z"/>`),
		g.Group(children),
	)
}