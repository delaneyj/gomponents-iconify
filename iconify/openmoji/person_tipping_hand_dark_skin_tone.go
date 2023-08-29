package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonTippingHandDarkSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#92d3f5" d="M11.934 60.943v-4.219c0-4.994 5.008-9 10-9q9 7.5 18 0c1.369 0 3.123-.278 4.39.261a11.877 11.877 0 0 1 3.236 2.813v10.145"/><path fill="#6a462f" d="M19.871 30.985c0 9 4.936 14 11 14c5.936 0 11.063-5 11.063-14a12.137 12.137 0 0 0-1-5c-3-3-7-8-7-8c-4 3-7 6-13 7c0 0-1.063 1-1.063 6Zm35.063 29.958L52.998 46.68l.938-1.25l6-2l5-6a1.414 1.414 0 0 0-2-2l-3 3c-1 1-4 0-7 1s-5 3.233-5 5v16.514"/><path d="M41.083 26.077c-2.661-.856-7.366-7.937-7.366-7.937c-2.662 3.232-12.737 6.986-12.737 6.986c-2.995 2.995-.046 13.86-.046 13.86c-4 0-4-6-4-13s4-14 14-14s14 7 14 14s0 13-4 13c1.325-2.253.15-12.91.15-12.91Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.934 38.985c-4 0-4-6-4-13s4-14 14-14s14 7 14 14s0 13-4 13m-29 21v-3c0-4.994 5.008-9 10-9q9 7.5 18 0a10.271 10.271 0 0 1 4.003.84"/><path d="M36.807 29.985a2 2 0 1 1-2-2a2 2 0 0 1 2 2m-7.999 0a2 2 0 1 1-2-2a2 2 0 0 1 2 2"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M19.871 30.985c0 9 4.936 14 11 14c5.936 0 11.063-5 11.063-14a12.137 12.137 0 0 0-1-5c-3-3-7-8-7-8c-4 3-7 6-13 7c0 0-1.063 1-1.063 6Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M27.934 37.985a6.553 6.553 0 0 0 6 0m21 22L52.998 46.68l.938-1.25l6-2l5-6a1.414 1.414 0 0 0-2-2l-3 3c-1 1-4 0-7 1s-5 3.233-5 5v15.556"/>`),
		g.Group(children),
	)
}