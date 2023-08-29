package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Noentry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="255.107" cy="256" r="254.106" fill="#FF473E"/><path fill="#D32A2A" d="M256 487.106C119.516 487.106 8.18 379.5 2.165 244.5c-.17 3.813-.271 7.645-.271 11.5c0 140.339 113.767 254.106 254.106 254.106S510.106 396.339 510.106 256c0-3.855-.101-7.687-.271-11.5C503.82 379.5 392.484 487.106 256 487.106z"/><path fill="#FFF" d="M415.997 311.61H94.216c-9.262 0-16.771-7.509-16.771-16.771v-82.431c0-9.262 7.509-16.771 16.771-16.771h321.781c9.262 0 16.771 7.509 16.771 16.771v82.431c0 9.262-7.508 16.771-16.771 16.771z"/>`),
		g.Group(children),
	)
}