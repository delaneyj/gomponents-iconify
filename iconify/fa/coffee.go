package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1664 384q0-80-56-136t-136-56h-64v384h64q80 0 136-56t56-136zM0 1152h1792q0 106-75 181t-181 75H256q-106 0-181-75T0 1152zm1856-768q0 159-112.5 271.5T1472 768h-64v32q0 92-66 158t-158 66H480q-92 0-158-66t-66-158V64q0-26 19-45t45-19h1152q159 0 271.5 112.5T1856 384z"/>`),
		g.Group(children),
	)
}