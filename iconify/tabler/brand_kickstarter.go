package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandKickstarter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 9l2.975-4.65C14.59 3.45 15.38 3 16.352 3c.79 0 1.474.286 2.054.858c.576.574.866 1.256.866 2.054c0 .588-.153 1.109-.46 1.559L16 11.5l3.465 4.912c.356.46.535 1 .535 1.613a2.92 2.92 0 0 1-.843 2.098c-.561.584-1.242.877-2.04.877c-.876 0-1.545-.29-2-.87l-4.112-5.697V17.5c0 .876-.313 1.69-.611 2.175C9.851 20.558 9.044 21 8.005 21c-.944 0-1.753-.327-2.271-.974c-.486-.6-.729-1.392-.729-2.38V6.275c0-.934.247-1.706.74-2.313C6.257 3.321 7.092 3 8.005 3c.868 0 1.821.321 2.4.962c.323.356.515.714.6 1.08c.052.224 0 .643 0 1.26V9z"/>`),
		g.Group(children),
	)
}