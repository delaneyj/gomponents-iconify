package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircledM(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#5092FF" d="M16 30c7.732 0 14-6.268 14-14S23.732 2 16 2S2 8.268 2 16s6.268 14 14 14Z"/><path fill="#fff" d="M22.438 9.564c.488.225.8.713.8 1.25v10.76a1.375 1.375 0 1 1-2.75 0v-7.499a.125.125 0 0 0-.206-.095l-3.388 2.905a1.375 1.375 0 0 1-1.757.027l-3.422-2.756a.125.125 0 0 0-.203.098v7.32a1.375 1.375 0 1 1-2.75 0V11.12a1.375 1.375 0 0 1 2.237-1.071l4.892 3.939a.125.125 0 0 0 .16-.002l4.918-4.216c.407-.35.981-.43 1.469-.206Z"/></g>`),
		g.Group(children),
	)
}