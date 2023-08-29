package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#F70000"/><path fill="#FFF" d="M15.947 13.347L5 24.89L15.996 7L27 25L15.947 13.347zm1.588 4.585h-3.422l1.76-1.936l1.662 1.953v-.017zm-6.6 3.177v-.024l1.941-1.987v-.009l5.92-.025l1.998 2.045h-9.858z"/></g>`),
		g.Group(children),
	)
}