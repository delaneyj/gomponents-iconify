package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chalkboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M96 32c-35.3 0-64 28.7-64 64v288h64V96h384v288h64V96c0-35.3-28.7-64-64-64H96zm128 352v32H32c-17.7 0-32 14.3-32 32s14.3 32 32 32h512c17.7 0 32-14.3 32-32s-14.3-32-32-32H416v-32c0-17.7-14.3-32-32-32H256c-17.7 0-32 14.3-32 32z"/>`),
		g.Group(children),
	)
}