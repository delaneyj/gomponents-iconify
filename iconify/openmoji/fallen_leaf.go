package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FallenLeaf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#a57939" d="M55.732 36.689c5.432 6.231 4.132 16.317 4.132 16.317s-9.931-.223-15.363-6.454s-4.132-16.317-4.132-16.317s10.046.069 15.363 6.454Z"/><path fill="#e27022" d="M19.762 22.361C26.154 17.008 36.05 18.5 36.05 18.5s-.47 10.065-6.762 15.254S13 37.616 13 37.616s.37-9.9 6.762-15.255Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.762 22.361C26.154 17.008 36.05 18.5 36.05 18.5s-.47 10.065-6.762 15.254S13 37.616 13 37.616s.37-9.9 6.762-15.255Zm35.97 14.328c5.432 6.231 4.132 16.317 4.132 16.317s-9.931-.223-15.363-6.454s-4.132-16.317-4.132-16.317s10.046.069 15.363 6.454Z"/>`),
		g.Group(children),
	)
}