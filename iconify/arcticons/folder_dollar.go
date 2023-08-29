package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderDollar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.5a3 3 0 0 1 3-3h8.718a4 4 0 0 1 2.325.745l4.914 3.51a4 4 0 0 0 2.325.745H40.5a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3h-33a3 3 0 0 1-3-3v-25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.014 31.497c.784 1.022 1.769 1.403 3.138 1.403h1.895a3.197 3.197 0 0 0 3.193-3.2h0c0-1.767-1.43-3.2-3.193-3.2h-2.094a3.197 3.197 0 0 1-3.193-3.2h0c0-1.767 1.43-3.2 3.193-3.2h1.895c1.37 0 2.354.38 3.138 1.403M24 34.5v-16"/>`),
		g.Group(children),
	)
}