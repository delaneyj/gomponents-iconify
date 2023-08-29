package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlChan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.794 18.907v10.186h5.063m-.023-4.376H19.54m-4.515 4.376l-3.291-10.186l-3.418 10.186m1.139-3.438h4.431m13.824-4.664v5.933m0-2.448a1.496 1.496 0 0 1 1.5-1.483h0a1.496 1.496 0 0 1 1.5 1.483v2.448m-4.487-.742a1.454 1.454 0 0 1-1.275.742h0a1.496 1.496 0 0 1-1.5-1.483v-.964a1.496 1.496 0 0 1 1.5-1.484h0a1.454 1.454 0 0 1 1.275.742m8.974 1.705a1.496 1.496 0 0 1-1.5 1.484h0a1.496 1.496 0 0 1-1.5-1.483v-.965a1.496 1.496 0 0 1 1.5-1.483h0a1.496 1.496 0 0 1 1.5 1.483m0 2.448v-3.931m4.487 3.931v-2.448a1.496 1.496 0 0 0-1.5-1.483h0a1.496 1.496 0 0 0-1.5 1.483v2.448m.001-2.448v-1.483"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.48 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33.04a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}