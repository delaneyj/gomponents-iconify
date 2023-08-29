package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whatsnew(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 9.5v29m39-29v29m-7.8-29v29m-7.8-29v29m-7.8-29v29m-7.8-29v29"/>`),
		g.Group(children),
	)
}