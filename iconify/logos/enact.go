package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Enact(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="#5582FF"><path d="M127.948 248.133L50.31 197.409v-28.675l77.742 50.879l-.104-64.284l-77.638-45.238V83.28L128 126.913l.104.052l-.156-64.388L0 .052v219.51l128.052 92.648z"/><path d="M128.052 312.21L256 219.665v-55.123l-128.052 83.591zm0-92.597l89.75-58.694V103l-89.854 52.329zm-.052-92.7l.104.052L256 55.072V0L127.948 62.577z" opacity=".8"/></g>`),
		g.Group(children),
	)
}