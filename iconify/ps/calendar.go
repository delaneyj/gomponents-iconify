package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 43h-43V21q0-21-21-21t-21 21v22H149V21q0-21-21-21t-21 21v22H64q-27 0-45.5 19.5T0 109v358q0 19 12.5 32T43 512h426q18 0 30.5-13t12.5-32V109q0-27-18.5-46.5T448 43zm21 426H43V171h426v298zm0-341H43v-19q0-10 6-17t15-7h384q9 0 15 7t6 17v19z"/>`),
		g.Group(children),
	)
}