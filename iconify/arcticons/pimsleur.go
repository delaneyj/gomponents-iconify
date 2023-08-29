package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pimsleur(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.013 20.513a16.013 16.013 0 0 0-32.026 0c0 8.844 6.572 16.992 16.013 22.987c0-7.389 16.013-8.51 16.013-22.987Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.843 8.73A24.327 24.327 0 0 0 19.85 30.785A24.896 24.896 0 0 0 24 43.5c.112-7.949 5.507-18.322 15.987-22.023"/>`),
		g.Group(children),
	)
}