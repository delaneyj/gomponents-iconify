package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sodexo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1m28.988 1.79l-3.94 5.22m3.94 0l-3.94-5.22"/><rect width="3.94" height="5.221" x="13.36" y="21.39" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.97"/><rect width="3.94" height="5.221" x="36.06" y="21.39" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.97"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.124 26.17a2.216 2.216 0 0 0 1.62.44h.442a1.304 1.304 0 0 0 1.302-1.305h0A1.304 1.304 0 0 0 10.186 24h-.884A1.304 1.304 0 0 1 8 22.695h0a1.304 1.304 0 0 1 1.302-1.305h.442a2.216 2.216 0 0 1 1.62.44m17.478 3.786a1.97 1.97 0 0 1-1.712.994h0a1.97 1.97 0 0 1-1.97-1.97v-1.28a1.97 1.97 0 0 1 1.97-1.97h0a1.97 1.97 0 0 1 1.97 1.97V24h-3.94m-1.99-.64a1.97 1.97 0 0 0-1.97-1.97h0a1.97 1.97 0 0 0-1.97 1.97v1.28a1.97 1.97 0 0 0 1.97 1.97h0a1.97 1.97 0 0 0 1.97-1.97m.001 1.97v-7.88"/>`),
		g.Group(children),
	)
}