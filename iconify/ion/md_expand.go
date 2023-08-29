package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdExpand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M396.795 396.8H320V448h128V320h-51.205z" fill="currentColor"/><path d="M396.8 115.205V192H448V64H320v51.205z" fill="currentColor"/><path d="M115.205 115.2H192V64H64v128h51.205z" fill="currentColor"/><path d="M115.2 396.795V320H64v128h128v-51.205z" fill="currentColor"/>`),
		g.Group(children),
	)
}