package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Noxcleaner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.532 42.5v-.002a31.896 31.896 0 0 1-1.015-7.978V23.803a3.672 3.672 0 0 0-3.672-3.672h-7.242V9.307c0-1.951-1.461-3.7-3.41-3.802a3.604 3.604 0 0 0-3.796 3.598v11.028h-7.242a3.672 3.672 0 0 0-3.672 3.672V34.52a31.88 31.88 0 0 1-1.015 7.98h31.063Zm-1.015-13.516H9.483M24 42.5v-3.097M34.145 42.5v-3.097M13.855 42.5v-3.097"/>`),
		g.Group(children),
	)
}