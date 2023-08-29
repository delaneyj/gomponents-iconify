package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Auspost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.753 31.973v13.41m0-42.765v2.474M16.296 3.93v40.143m5.457-38.98a13.514 13.514 0 1 1 0 26.88m0-29.354a21.5 21.5 0 1 1 0 42.764m-5.457-1.31a21.5 21.5 0 0 1 0-40.145"/>`),
		g.Group(children),
	)
}