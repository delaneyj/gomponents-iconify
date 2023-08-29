package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BitbucketLogoThirtyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1981 64q28 0 47 19t20 47q0 7-1 11l-278 1723q-4 25-22 40t-44 16H367q-32 0-57-21t-31-54L1 141q-1-4-1-12q0-28 19-46t48-19h1914zm-638 624H697l116 608h426l104-608z"/>`),
		g.Group(children),
	)
}