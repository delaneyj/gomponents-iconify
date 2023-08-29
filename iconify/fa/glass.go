package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1635 58q0 35-43 78L960 768v768h320q26 0 45 19t19 45t-19 45t-45 19H384q-26 0-45-19t-19-45t19-45t45-19h320V768L72 136Q29 93 29 58q0-23 18-36.5T85 4t43-4h1408q23 0 43 4t38 17.5t18 36.5z"/>`),
		g.Group(children),
	)
}