package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rapido(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 4.877a1.237 1.237 0 1 1 .029 1.749a1.237 1.237 0 0 1-.03-1.75Zm25 0a1.237 1.237 0 1 1-1.748-.03a1.237 1.237 0 0 1 1.749.03ZM16.74 34.523c-.107-11.532 14.26-11.3 14.62 0m-14.62 0V15.658m0 0h14.62m0 0v18.865m-7.552-5.901l.006 14.878"/><circle cx="23.876" cy="10.209" r="2.229" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}