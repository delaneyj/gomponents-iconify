package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tapet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 17.51h6.429m-3.214 9.705V17.51m24.511 1.274v8.431m-1.273-6.43H38.5m-18.348 4.003a2.426 2.426 0 0 1-2.426 2.427h0a2.426 2.426 0 0 1-2.426-2.426V23.21a2.426 2.426 0 0 1 2.426-2.426h0a2.426 2.426 0 0 1 2.426 2.426m0 4.005v-6.43m14.154 5.205a2.425 2.425 0 0 1-2.108 1.225h0a2.426 2.426 0 0 1-2.426-2.427v-1.577a2.426 2.426 0 0 1 2.426-2.426h0a2.426 2.426 0 0 1 2.426 2.426V24h-4.852m-7.16.788a2.426 2.426 0 0 0 2.426 2.427h0a2.426 2.426 0 0 0 2.426-2.426V23.21a2.426 2.426 0 0 0-2.426-2.426h0a2.426 2.426 0 0 0-2.426 2.426m0-2.425v9.705"/>`),
		g.Group(children),
	)
}