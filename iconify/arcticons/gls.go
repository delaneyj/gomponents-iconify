package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.71 17.33v12.734h6.63m2.713-1.238a3.724 3.724 0 0 0 3.14 1.432h1.885a3.182 3.182 0 0 0 0-6.363h-2.041a3.182 3.182 0 0 1 0-6.363h1.884a3.354 3.354 0 0 1 3.14 1.432m-20.523 2.739a4.282 4.282 0 0 0-4.397-4.295a4.46 4.46 0 0 0-3.925 4.454v3.978a4.24 4.24 0 1 0 8.479 0h-4.24"/><circle cx="39.034" cy="29.919" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}