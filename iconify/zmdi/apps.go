package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 107V21h85v86H0zm128 256v-86h85v86h-85zM0 363v-86h85v86H0zm0-128v-86h85v86H0zm128 0v-86h85v86h-85zM256 21h85v86h-85V21zm-128 86V21h85v86h-85zm128 128v-86h85v86h-85zm0 128v-86h85v86h-85z"/>`),
		g.Group(children),
	)
}