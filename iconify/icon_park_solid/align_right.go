package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAlignRight0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M6 17h28v14H6z"/><path stroke-linecap="round" d="M42 6v36"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAlignRight0)"/>`),
		g.Group(children),
	)
}