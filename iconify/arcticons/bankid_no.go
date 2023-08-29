package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankidNo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Zm-29 10.6h5.1m-5.1 15.8h5.1m-5.1-5.3h5.1m4.8-5.2h5.2m-5.2 5.2h5.2m4.8-10.5h5.1m-5.1 5.3h5.1m-5.1 10.5h5.1"/>`),
		g.Group(children),
	)
}