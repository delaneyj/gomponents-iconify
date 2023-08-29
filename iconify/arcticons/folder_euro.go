package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderEuro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.5a3 3 0 0 1 3-3h8.718a4 4 0 0 1 2.325.745l4.914 3.51a4 4 0 0 0 2.325.745H40.5a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3h-33a3 3 0 0 1-3-3v-25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.01 32.588a5.354 5.354 0 0 1-4.107 1.912h0a5.365 5.365 0 0 1-5.365-5.365v-5.27a5.365 5.365 0 0 1 5.365-5.365h0c1.658 0 3.14.752 4.124 1.933m-12.054 4.123h6.775m-6.775 3.926h6.775"/>`),
		g.Group(children),
	)
}