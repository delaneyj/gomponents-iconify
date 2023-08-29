package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EasyJet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.523 19.249v5.328a1.774 1.774 0 0 1-1.773 1.776h0a1.774 1.774 0 0 1-1.772-1.776v-.59m-1.624.59v2.398a1.776 1.776 0 0 1-1.776 1.776h0a1.77 1.77 0 0 1-1.256-.52"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.354 21.646v2.931a1.776 1.776 0 0 1-1.776 1.776h0a1.776 1.776 0 0 1-1.776-1.776v-2.93m-4.655 4.309a1.998 1.998 0 0 0 1.461.397h.398a1.176 1.176 0 0 0 1.175-1.176h0A1.175 1.175 0 0 0 20.006 24h-.796a1.175 1.175 0 0 1-1.174-1.177h0a1.175 1.175 0 0 1 1.174-1.177h.399a1.998 1.998 0 0 1 1.46.398m-9.749 3.413a1.775 1.775 0 0 1-1.543.896h0A1.776 1.776 0 0 1 8 24.578v-1.154a1.776 1.776 0 0 1 1.776-1.776h0a1.776 1.776 0 0 1 1.776 1.776V24H8m28.493 1.457a1.776 1.776 0 0 1-1.543.896h0a1.776 1.776 0 0 1-1.776-1.776v-1.154a1.776 1.776 0 0 1 1.776-1.776h0a1.776 1.776 0 0 1 1.776 1.776V24h-3.552m5.672-3.819v5.284a.888.888 0 0 0 .888.889H40m-2.087-4.708h1.865M16.54 24.577a1.776 1.776 0 0 1-1.776 1.776h0a1.776 1.776 0 0 1-1.776-1.776v-1.154a1.776 1.776 0 0 1 1.776-1.776h0a1.776 1.776 0 0 1 1.776 1.776m0 2.931v-4.708"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`),
		g.Group(children),
	)
}