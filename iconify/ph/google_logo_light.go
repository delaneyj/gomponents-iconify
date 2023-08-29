package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M222 128a94 94 0 1 1-21.49-59.82a6 6 0 1 1-9.25 7.64A82 82 0 1 0 209.78 134H128a6 6 0 0 1 0-12h88a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}