package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Npm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M71.609 112.569v286.649h143.432v-215.04h71.608v215.04h71.608V112.569H71.608zM430.08 40.96v430.08H0V40.96h430.08z"/>`),
		g.Group(children),
	)
}