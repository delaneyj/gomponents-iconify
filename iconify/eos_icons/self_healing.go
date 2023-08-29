package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelfHealing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.048 11.769A8.187 8.187 0 0 1 5.73 5.722L8.167 8.17V2H2l2.276 2.262A10.271 10.271 0 0 0 10.214 22v-2.077a8.218 8.218 0 0 1-7.166-8.154ZM23 12.283A10.315 10.315 0 0 0 13.786 2v2.077a8.27 8.27 0 0 1 7.166 8.206a8.074 8.074 0 0 1-2.682 5.995l-2.484-2.448V22h6.19l-2.252-2.262A10.12 10.12 0 0 0 23 12.283Z"/><path fill="currentColor" d="m12.29 16l-.537-.49c-1.91-1.732-3.172-2.875-3.172-4.277a2.02 2.02 0 0 1 2.04-2.04a2.221 2.221 0 0 1 1.67.775a2.221 2.221 0 0 1 1.669-.775a2.02 2.02 0 0 1 2.04 2.04c0 1.402-1.261 2.545-3.172 4.281Z"/>`),
		g.Group(children),
	)
}