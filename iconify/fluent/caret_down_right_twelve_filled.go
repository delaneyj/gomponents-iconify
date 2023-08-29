package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDownRightTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.005 3.963c0-.891-1.077-1.337-1.707-.707L3.255 7.299c-.63.63-.184 1.707.707 1.707h3.543a1.5 1.5 0 0 0 1.5-1.5V3.963Z"/>`),
		g.Group(children),
	)
}