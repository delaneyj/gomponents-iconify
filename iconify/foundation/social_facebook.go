package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialFacebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M82.026 14H17.974A3.974 3.974 0 0 0 14 17.974v64.053A3.974 3.974 0 0 0 17.974 86h34.483V58.118h-9.383V47.252h9.383v-8.014c0-9.3 5.68-14.363 13.976-14.363c3.974 0 7.389.295 8.385.428v9.719l-5.754.003c-4.512 0-5.385 2.144-5.385 5.29v6.938h10.76l-1.401 10.866h-9.359V86h18.348A3.974 3.974 0 0 0 86 82.026V17.974A3.974 3.974 0 0 0 82.026 14z"/>`),
		g.Group(children),
	)
}