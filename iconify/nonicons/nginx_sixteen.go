package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NginxSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.62.102a.757.757 0 0 1 .76 0l6.246 3.625a.75.75 0 0 1 .374.648v7.25a.75.75 0 0 1-.374.648L8.38 15.898a.757.757 0 0 1-.76 0l-6.246-3.625A.75.75 0 0 1 1 11.625v-7.25a.75.75 0 0 1 .374-.648L7.62.102ZM2.508 4.806v6.388L8 14.382l5.492-3.188V4.806L8 1.618L2.508 4.806Zm2.475-.249a.757.757 0 0 1 .822.163l4.241 4.22V5.25c0-.414.338-.75.754-.75s.754.336.754.75v5.5a.75.75 0 0 1-.466.693a.757.757 0 0 1-.821-.163L6.026 7.06v3.69c0 .414-.338.75-.754.75a.752.752 0 0 1-.754-.75v-5.5a.75.75 0 0 1 .465-.693Z"/>`),
		g.Group(children),
	)
}