package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadsetOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 14v-3c0-1.953.7-3.742 1.862-5.13m2.182-1.825A8 8 0 0 1 20 11v3m-2 5c0 1.657-2.686 3-6 3"/><path d="M4 14a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-3zm12.169-1.82c.253-.115.534-.18.831-.18h1a2 2 0 0 1 2 2v2m-1.183 2.826c-.25.112-.526.174-.817.174h-1a2 2 0 0 1-2-2v-2M3 3l18 18"/></g>`),
		g.Group(children),
	)
}