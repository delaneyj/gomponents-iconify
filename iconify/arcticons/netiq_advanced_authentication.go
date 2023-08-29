package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetiqAdvancedAuthentication(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 6.433H22.267A2.567 2.567 0 0 0 19.7 8.999v8.342m8.324 18.608h-5.757a2.567 2.567 0 0 1-2.567-2.567v-8.341"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 35.949L29.819 41.46a1.283 1.283 0 0 1-1.795-1.177V14.406a2.567 2.567 0 0 1 1.544-2.353L42.5 6.433Z"/><circle cx="9.583" cy="21.191" r="4.083" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.123 25.273v-4.082h0h-11.462m8.783 2.354v-2.354"/>`),
		g.Group(children),
	)
}