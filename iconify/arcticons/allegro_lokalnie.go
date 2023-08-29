package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AllegroLokalnie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.835 21.628v8.469a5.321 5.321 0 0 1-5.37 5.369H29.37L24 40.834l-5.369-5.368h-5.733a5.524 5.524 0 0 1-5.733-5.37v-8.689m.968-11.979a5.332 5.332 0 0 1 4.401-2.262h22.932a5.388 5.388 0 0 1 4.419 2.322"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 25.145a2.639 2.639 0 1 1-2.639 2.639A2.64 2.64 0 0 1 24 25.144Zm-7.644 0a2.639 2.639 0 1 0 2.64 2.639a2.64 2.64 0 0 0-2.64-2.64Zm15.288 0a2.639 2.639 0 1 1-2.64 2.639a2.64 2.64 0 0 1 2.64-2.64Zm-11.492-5.142l.318-10.3m-7.716 10.3l1.248-10.348M24 9.563H7.598l-2.97 6.452a4.838 4.838 0 0 0 1.392 4.657c2.146 1.877 5.607 1.43 6.735-.669c0 0 3.549 3.849 7.397 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.848 20.003l-.318-10.3m7.716 10.3L33.998 9.655M24 9.563h16.402l2.97 6.451a4.838 4.838 0 0 1-1.392 4.658c-2.145 1.877-5.606 1.43-6.734-.669c0 0-3.55 3.849-7.398 0c0 0-3.817 3.88-7.697 0"/>`),
		g.Group(children),
	)
}