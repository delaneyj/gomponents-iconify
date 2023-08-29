package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonLightSkinToneBeard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#92D3F5" d="M17 61v-4c0-5 5-9 10-9c6 5 12 5 18 0c5 0 10 4 10 9v4"/><path fill="#fadcbc" d="M24.9 31c0 9 4.9 14 11 14c6 0 11.1-5 11.1-14a12.408 12.408 0 0 0-1-5c-3-3-7-8-7-8c-4 3-7 6-13 7c0 0-1.1 1-1.1 6z"/><path fill="#a57939" d="M36 12c-10 0-14 7-14 14c0 6.7 0 12.4 3.5 12.9c.2.5.4.9.6 1.4a13.95 13.95 0 0 0 2.1 2.7a9.996 9.996 0 0 0 7.4 2.5h.3a9.996 9.996 0 0 0 7.4-2.5a8.972 8.972 0 0 0 2.1-2.7c.374-.714.708-1.448 1-2.2l-.3.9c3.9-.1 3.9-6 3.9-13s-4-14-14-14zm4 27c0 .7-1.6 2-4 2s-4-1.3-4-2v-1.9a1.002 1.002 0 0 1 1-1l3 1.5l3.1-1.5a.945.945 0 0 1 1 1L40 39zm4.9-4.2s-1.9 3-2.7-1.2a24.285 24.285 0 0 0-13.1 0c-.7 4.2-2.7 1.2-2.7 1.2a29.028 29.028 0 0 1-1.6-4.1c.192-1.864.56-3.705 1.1-5.5c7.2-1 12.9-7.1 12.9-7.1s1.4 1.2 7.2 7.8a33.87 33.87 0 0 1 .7 4a19.05 19.05 0 0 1-1.8 4.9z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M25.5 39.2C20.7 39.2 22 33 22 26s4-14 14-14s14 7 14 14s.9 13.4-4 13M17 60v-3c0-5 5-9 10-9c6 5 12 5 18 0c5 0 10 4 10 9v3"/><path d="M41.9 29a2 2 0 1 1-4 0a2 2 0 0 1 4 0m-8 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M47 31a12.408 12.408 0 0 0-1-5c-3-3-7-8-7-8c-4 3-7 6-13 7c0 0-1.1 1-1.1 6"/><path d="M39.9 36.5a.992.992 0 0 1-.5 1.3a7.19 7.19 0 0 1-3.5.8a8.075 8.075 0 0 1-3.4-.8a.992.992 0 0 1-.5-1.3a.883.883 0 0 1 1.2-.5c.1 0 .1 0 .1.1c1.58.9 3.52.9 5.1 0a1.129 1.129 0 0 1 1.5.4z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M24.9 31c-.2 3.1-.3 6.6 1.2 9.3a13.95 13.95 0 0 0 2.1 2.7a10.582 10.582 0 0 0 7.4 2.5"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M47 31a8.389 8.389 0 0 1-1.9 3.9s-1.9 3-2.7-1.2a24.285 24.285 0 0 0-13.1 0c-.8 4.2-2.7 1.2-2.7 1.2A11.1 11.1 0 0 1 25 31"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M47 31.4a16.113 16.113 0 0 1-1.6 8.9a8.972 8.972 0 0 1-2.1 2.7a9.996 9.996 0 0 1-7.4 2.5h-.3"/>`),
		g.Group(children),
	)
}