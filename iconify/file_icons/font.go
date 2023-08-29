package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Font(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 256h192v64h-64v192H96V320H32zm448-128H354.125v384h-68.25V128H160V64h320z"/>`),
		g.Group(children),
	)
}