package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClusterRoleBinding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 10h5v2H5Zm3.54 12H4a2.006 2.006 0 0 1-2-2V6a2.006 2.006 0 0 1 2-2h4.18a2.988 2.988 0 0 1 5.64 0H18a2.006 2.006 0 0 1 2 2v6.09a5.989 5.989 0 0 0-1-.09h-1V9H4v4h5.69a6.04 6.04 0 0 0-1.878 2H4v4h3.09a5.973 5.973 0 0 0 1.45 3ZM10 5a1 1 0 1 0 1-1a1.003 1.003 0 0 0-1 1Zm7 6a1 1 0 1 0-1 1a1 1 0 0 0 1-1ZM5 18h2a5.96 5.96 0 0 1 .35-2H5Zm9-7a1 1 0 1 0-1 1a1 1 0 0 0 1-1Zm1 9h-2a2 2 0 0 1 0-4h2v-2h-2a4 4 0 0 0 0 8h2Zm8-2a4 4 0 0 1-4 4h-2v-2h2a2 2 0 0 0 0-4h-2v-2h2a4 4 0 0 1 4 4Zm-5 1h-4v-2h4Z"/>`),
		g.Group(children),
	)
}