package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zM7 11.09v11.974l9-5.315l9 5.315V11.09l-2.606 1.13v6.353L16 14.786l-6.333 3.787V12.22L7 11.09zm.303-.489l2.364.978L16 7.883l6.394 3.696l2.364-.978L16 5.5l-8.697 5.101zm1.879 11.821L16.03 26.5l6.849-4.078l-2.03-1.13l-4.819 2.825l-4.878-2.825l-1.97 1.13z"/>`),
		g.Group(children),
	)
}