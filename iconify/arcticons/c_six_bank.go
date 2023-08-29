package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CSixBank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.147 26.634v.066a5.3 5.3 0 0 1-5.3 5.3h0a5.3 5.3 0 0 1-5.3-5.3v-5.4a5.3 5.3 0 0 1 5.3-5.3h0a5.3 5.3 0 0 1 5.3 5.3v.066"/><circle cx="31.153" cy="26.7" r="5.3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.984 17.951A5.176 5.176 0 0 0 31.524 16h-.371a5.3 5.3 0 0 0-5.3 5.3v5.4"/>`),
		g.Group(children),
	)
}