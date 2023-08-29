package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PurchasedApps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.038 16.341h31.925a3.537 3.537 0 0 1 3.537 3.537v18.986a3.538 3.538 0 0 1-3.538 3.538H8.037A3.537 3.537 0 0 1 4.5 38.865V19.88a3.538 3.538 0 0 1 3.538-3.538Z"/><rect width="4.126" height="4.127" x="19.874" y="19.353" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.063"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.367 19.353h0a2.063 2.063 0 0 1 2.063 2.063h0a2.063 2.063 0 0 1-2.063 2.064h0a2.063 2.063 0 0 1-2.063-2.064h0a2.063 2.063 0 0 1 2.063-2.063Zm-13.473-7.141a6.612 6.612 0 1 1 13.224 0m-13.224 0v4.15m13.224-4.15v3.913m-22.487-3.331a6.612 6.612 0 0 1 10.432-5.4m-10.432 5.4v3.262M8.13 21.842a2.063 2.063 0 1 1 4.126 0m-.001 14.897a2.063 2.063 0 1 1-4.126 0m4.126 0V21.84m-4.125.002V36.74"/>`),
		g.Group(children),
	)
}