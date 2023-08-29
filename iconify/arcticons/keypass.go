package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keypass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.462 6.553h4.814L14.027 4.5h19.258l2.751 2.053h5.502V22.29c0 15.731-17.882 21.21-17.882 21.21S6.462 38.066 6.462 22.29Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.778 18.868h13.756v7.527c0 3.225-3.636 6.842-6.878 6.842s-6.878-3.617-6.878-6.842Zm6.878 4.79v3.474"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.842 18.868v-2.79a5.402 5.402 0 0 1 4.814-4.788a5.402 5.402 0 0 1 4.815 4.79v2.79"/>`),
		g.Group(children),
	)
}