package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aimsicd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4v4.83h4.26L24 16.21v-4.84h-4.26Zm0 15.07L37.22 42l-3.51 2L26 30.7V44h-4V30.7L14.29 44l-3.51-2L24 19.07Z"/>`),
		g.Group(children),
	)
}