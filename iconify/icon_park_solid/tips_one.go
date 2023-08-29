package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TipsOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTipsOne0"><g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M44 8H4v30h15l5 5l5-5h15V8Z"/><path stroke="#000" d="M24 23v9m0-16v1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTipsOne0)"/>`),
		g.Group(children),
	)
}