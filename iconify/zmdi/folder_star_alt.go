package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderStarAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 64q18 0 30.5 12.5T427 107v213q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21h128l42 43h171zM247 320l-20-87l67-58l-89-8l-34-82l-35 82l-89 8l68 58l-21 87l77-45z"/>`),
		g.Group(children),
	)
}