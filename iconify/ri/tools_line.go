package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.329 3.272a3.5 3.5 0 0 1 4.255 4.962l10.709 10.71l-1.415 1.414L8.17 9.648a3.502 3.502 0 0 1-4.962-4.255L5.443 7.63a1.5 1.5 0 0 0 2.122-2.121L5.329 3.272Zm10.367 1.883l3.182-1.768l1.415 1.415l-1.768 3.182l-1.768.353l-2.121 2.121l-1.415-1.414l2.122-2.121l.353-1.768ZM8.98 13.287l1.414 1.414l-5.303 5.303a1 1 0 0 1-1.492-1.327l.077-.087l5.304-5.303Z"/>`),
		g.Group(children),
	)
}