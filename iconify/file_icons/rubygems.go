package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rubygems(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m134 168l-63 64l153 153l153-153l-63-64H134zM223 0L0 128v256l223 128l224-128V128L223 0zm181 360L223 464L43 360V152L223 48l181 104v208z"/>`),
		g.Group(children),
	)
}