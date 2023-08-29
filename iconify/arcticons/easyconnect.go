package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Easyconnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.914 5.664H15.2l-9.7 9.81l6.908 6.76H34.82l-7.055-6.76Zm-9.148 9.811H5.5m8.12 10.143h22.34l6.54 6.724l-9.406 9.994h-22.89l9.994-9.994Zm6.576 6.723h22.302"/>`),
		g.Group(children),
	)
}