package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosPhotosOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M96 128v320h384V128H96zm368 304H112V144h352v288z" fill="currentColor"/><path d="M32 64v320h48v-16H48V80h352v32h16V64z" fill="currentColor"/>`),
		g.Group(children),
	)
}