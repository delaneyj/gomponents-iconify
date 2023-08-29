package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 299h299v42H0v-42zm96-90l-19 47H32L133 21h32l102 235h-45l-19-47H96zm53-145l-40 107h80z"/>`),
		g.Group(children),
	)
}