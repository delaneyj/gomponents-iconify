package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindPower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 28V15.816a3.006 3.006 0 0 0 1.47-1.119L27.063 17l.518-1.932l-8.605-2.305a2.985 2.985 0 0 0-2.63-2.728L14.194 2l-1.932.519l2.132 7.956a2.959 2.959 0 0 0-1.093 3.811L7 20.586L8.415 22l6.3-6.3a2.894 2.894 0 0 0 .285.115V28H2v2h28v-2Zm0-15a1 1 0 1 1-1-1a1 1 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}