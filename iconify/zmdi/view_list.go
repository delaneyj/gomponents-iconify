package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 235v-86h85v86H0zm0 106v-85h85v85H0zm0-213V43h85v85H0zm107 107v-86h256v86H107zm0 106v-85h256v85H107zm0-298h256v85H107V43z"/>`),
		g.Group(children),
	)
}