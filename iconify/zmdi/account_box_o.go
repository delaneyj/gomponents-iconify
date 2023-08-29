package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountBoxO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M192 197q-20 0-34-14t-14-34t14-34t34-14t34 14t14 34t-14 34t-34 14zm96 86v16H96v-16q0-22 33-35t63-13t63 13t33 35zM341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h298zm0 341V43H43v298h298z"/>`),
		g.Group(children),
	)
}