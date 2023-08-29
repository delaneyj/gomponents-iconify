package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 146.484v168.677l600 342.114l600-342.114V146.484H0zm0 276.563v494.604L305.64 597.29L0 423.047zm1200 0L894.36 597.29L1200 917.651V423.047zM389.575 645.19L0 1053.516h1200L810.425 645.19L600 765.161L389.575 645.19z"/>`),
		g.Group(children),
	)
}