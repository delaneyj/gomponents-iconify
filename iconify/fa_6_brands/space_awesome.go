package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpaceAwesome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M96 256h32v256H0V352h32v-32h32v-32h32v-32zm416 96v160H384V256h32v32h32v32h32v32h32zM320 64h32v384h-32v-32H192v32h-32V64h32V32h32V0h64v32h32v32zm-32 64h-64v64h64v-64z"/>`),
		g.Group(children),
	)
}