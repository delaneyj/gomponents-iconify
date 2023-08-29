package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.85 22.365l.993-9.808s1.167 6.83 2.276 7.94c.584-.175 2.628-.876 2.628-.876c1.167-3.036 2.277-12.26 2.277-12.26s.992 9.283 1.985 11.327c.642-.117 3.152-.35 3.152-.35c.934-2.395 1.927-13.838 1.927-13.838s1.226 11.385 2.101 13.837c.701.058 2.978.409 2.978.409c1.05-2.277 2.218-11.385 2.218-11.385s1.051 10.392 2.044 12.377c.817.175 2.569.817 2.569.817c1.284-1.518 2.393-7.765 2.393-7.765l.76 9.634S26.918 26.686 24.29 43.5C21.928 26.773 8.85 22.365 8.85 22.365Z"/>`),
		g.Group(children),
	)
}