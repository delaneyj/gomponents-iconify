package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Friendfeed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M462 81h-68q-10 0-19 12.5t-9 29.5v40h81v81h-81v203h-81V244H163v203H83V244H2v-81h81v-40q0-50 31.5-86T192 1h68v80h-68q-11 0-20 12.5t-9 29.5v40h122v-40q0-50 32-86t77-36h68v80z"/>`),
		g.Group(children),
	)
}