package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallBlockedTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M20.934 5.66c-1.96-1.73-5.021-2.593-9.185-2.588c-4.157.005-7.126.873-8.907 2.603c-.754.733-1.029 1.82-.721 2.851l.31 1.037c.288.967 1.261 1.613 2.276 1.51l2.035-.204c.873-.088 1.563-.71 1.711-1.547l.391-2.204a8.102 8.102 0 0 1 3.14-.708c1.126-.033 2.156.142 3.09.526l.632 2.355c.226.846.986 1.48 1.878 1.566l2.047.198c1.028.099 1.935-.553 2.12-1.525l.198-1.036c.195-1.028-.19-2.106-1.015-2.834z" fill="currentColor"/><path d="M12 22a5.5 5.5 0 1 1 0-11a5.5 5.5 0 0 1 0 11zm0-1.5a4 4 0 0 0 3.309-6.248l-5.557 5.557c.64.436 1.414.691 2.248.691zm-3.309-1.752l5.557-5.557a4 4 0 0 0-5.557 5.557z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}