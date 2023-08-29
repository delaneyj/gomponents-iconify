package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PolarisIqHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Zm-5.6 22l-5.8-5.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.1 21.5c0 3.3 2.7 6 5.8 6c3.3 0 6-2.7 6-6v-6c0-3.3-2.7-6-6-6s-5.8 2.7-5.8 6v6Zm-10 5.9h5.3M13.1 9.7h5.3m-2.6 0v17.7m3.6 11.4h0c-1 0-1.8-.8-1.8-1.8v-1.2c0-1 .8-1.8 1.8-1.8h0c1 0 1.8.8 1.8 1.8V37c.1 1-.8 1.8-1.8 1.8Zm4-2.9c0-1 .8-1.8 1.8-1.8h0c1 0 1.8.8 1.8 1.8v2.9m-3.6-4.7v4.7m3.7-2.9c0-1 .8-1.8 1.8-1.8h0c1 0 1.8.8 1.8 1.8v2.9m-19.2-7.3v7.3m0-3c0-1 .8-1.8 1.8-1.8h0c1 0 1.8.8 1.8 1.8v3m21.1-.9c-.3.5-.9.9-1.6.9h0c-1 0-1.8-.8-1.8-1.8v-1.2c0-1 .8-1.8 1.8-1.8h0c1 0 1.8.8 1.8 1.8v.6h-3.6"/>`),
		g.Group(children),
	)
}