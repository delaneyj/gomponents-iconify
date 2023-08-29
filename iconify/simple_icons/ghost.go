package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ghost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12S18.627 0 12 0zm.256 2.313c2.47.005 5.116 2.008 5.898 2.962l.244.3c1.64 1.994 3.569 4.34 3.569 6.966c0 3.719-2.98 5.808-6.158 7.508c-1.433.766-2.98 1.508-4.748 1.508c-4.543 0-8.366-3.569-8.366-8.112c0-.706.17-1.425.342-2.15c.122-.515.244-1.033.307-1.549c.548-4.539 2.967-6.795 8.422-7.408a4.29 4.29 0 0 1 .49-.026Z"/>`),
		g.Group(children),
	)
}