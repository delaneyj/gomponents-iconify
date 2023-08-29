package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amethyst(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M24 3.5L5.768 16.392c.046 5.072.021 10.144 0 15.216l3.687 2.608c.718-15.085 7.231-21.088 7.231-21.088s.198.607 2.385.858c2.188.25 9.393-1.17 11.418 1.095s-.38 4.21 1.31 5.895c1.689 1.685 6.252 2.196 5.747 3.74c-.505 1.545-3.782.645-7.836.636c-4.054-.008-3.43 1.352-6.61 1.27c-4.824 3.416-.468 11.369-3.412 14.817l-.012.004L24 44.5l18.23-12.892c-.045-5.072-.02-10.144 0-15.216L24 3.5Z"/><circle cx="26.819" cy="17.312" r="1.134"/></g>`),
		g.Group(children),
	)
}