package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#136AAD"/><path fill="#FFF" d="m7 11.09l2.667 1.13v6.353L16 14.786l6.394 3.787V12.22L25 11.09v11.974l-9-5.315l-9 5.315V11.09zm.303-.489L16 5.5l8.758 5.101l-2.364.978L16 7.883l-6.333 3.696l-2.364-.978zm1.879 11.821l1.97-1.13l4.878 2.825l4.818-2.825l2.03 1.13L16.03 26.5l-6.848-4.078z"/></g>`),
		g.Group(children),
	)
}