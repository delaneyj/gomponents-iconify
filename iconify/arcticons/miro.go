package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Miro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5h5.322l10.58 10.205L11.072 42.5H5.5l5.071-24.354L5.5 5.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.643 5.5h5.322l9.767 6.73l-9.517 30.27h-5.572l4.759-26.795L16.643 5.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.786 5.5h5.322l9.392 5.916L33.358 42.5h-5.572l3.946-30.27l-3.946-6.73z"/>`),
		g.Group(children),
	)
}