package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TipsOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTipsOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M44 8H4v30h15l5 5l5-5h15V8Z"/><path d="M24 23v9m0-16v1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTipsOne0)"/>`),
		g.Group(children),
	)
}