package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChildLightSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#92D3F5" d="M18 61v-4.1c0-5 3.8-9 8.8-9c6 5 12 5 18 0c5 0 9.2 4 9.2 9V61"/><path fill="#a57939" d="M26 40c-6 0-5-16-2-20c6-8 18-8 24 0c3 4 4 20-2 20"/><path fill="#fadcbc" d="M25 34c0 8 4.8 11.9 10.8 11.9S47 42 47 34c0-5.9-2.2-9.2-2.2-9.2c-2 0-5.8-3.5-5.8-3.5s-8.2 5.5-12.8 5.5c0 0-1.2 1.2-1.2 7.2z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M26 40c-6 0-5-16-2-20c6-8 18-8 24 0c3 4 4 20-2 20M18 60v-3.1c0-5 3.8-9 8.8-9c6 5 12 5 18 0c5 0 9.2 4 9.2 9V60"/><path d="M42 33c0 1.1-.9 2-2 2s-2-.9-2-2s.9-2 2-2s2 .9 2 2m-8 0c0 1.1-.9 2-2 2s-2-.9-2-2s.9-2 2-2s2 .9 2 2"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M25 34c0 8 4.8 11.9 10.8 11.9S47 42 47 34c0-5.9-2.2-9.2-2.2-9.2c-2 0-5.8-3.5-5.8-3.5s-8.2 5.5-12.8 5.5c0 0-1.2 1.2-1.2 7.2z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M32.8 39.2c1.9 1 4.1 1 6 0"/>`),
		g.Group(children),
	)
}