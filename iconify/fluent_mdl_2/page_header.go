package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageHeader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M127 0h1792v2048H127V0zm1665 1921V129H256v1792h1536zM1663 258v639H383V258h1280zm-127 510V387H512v381h1024zm127 256v641H383v-641h1280zm-127 512v-383H512v383h1024z"/>`),
		g.Group(children),
	)
}