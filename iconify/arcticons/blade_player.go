package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BladePlayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.958 17.269a10.208 10.208 0 0 0-12.831-2.937v-3.206L5.5 5.5v17.96a10.27 10.27 0 0 0 16.474 8.177"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.004 23.462l-3.312 1.91l-3.311 1.91V19.64l3.311 1.91Zm1.958 13.506V23.46a10.332 10.332 0 1 1 5.627 9.128V42.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.466 23.46l-3.311-1.91l-3.312-1.91v7.64l3.312-1.91Z"/>`),
		g.Group(children),
	)
}