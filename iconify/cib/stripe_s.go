package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StripeS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13.706 9.663c0-1.394 1.162-1.931 3.025-1.931c2.713 0 6.156.831 8.869 2.294V1.633C22.644.452 19.694.002 16.737.002c-7.231 0-12.05 3.775-12.05 10.087c0 9.869 13.55 8.269 13.55 12.525c0 1.65-1.431 2.181-3.419 2.181c-2.95 0-6.763-1.219-9.756-2.844v8.031a24.75 24.75 0 0 0 9.75 2.025c7.413 0 12.519-3.188 12.519-9.6c0-10.637-13.625-8.731-13.625-12.744z"/>`),
		g.Group(children),
	)
}