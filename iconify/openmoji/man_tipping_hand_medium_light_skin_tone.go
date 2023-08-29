package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManTippingHandMediumLightSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#92d3f5" d="M11.945 60.958V56.74c0-4.994 5.008-9 10-9q9 7.5 18 0c1.368 0 3.122-.278 4.389.26a11.877 11.877 0 0 1 3.236 2.813v10.145"/><path fill="#debb90" d="m54.945 60.958l-1.937-14.263l.937-1.25l6-2l5-6a1.414 1.414 0 0 0-2-2l-3 3c-1 1-4 0-7 1s-5 3.232-5 5v16.513M19.886 31c-.102 7.999 4.063 14 11 14c7.063 0 11.063-6 11.063-14c0-5-3-10-3-10c-8 0-10 3-16 1a17.077 17.077 0 0 0-3.063 9Z"/><path fill="#fcea2b" d="M39.227 19.688c-7.473 0-8.716 3.591-16.36 1.88C19.463 20.806 20.95 38 20.95 38c-3 0-4-7-4-14c0-6 5-12 14-12s14 6 14 12c0 7-1 14-4 14c0 0 .717-18.312-1.722-18.312Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.945 60v-3c0-4.994 5.008-9 10-9q9 7.5 18 0a10.271 10.271 0 0 1 4.003.84"/><path d="M36.818 30a2 2 0 1 1-2-2a2 2 0 0 1 2 2m-8 0a2 2 0 1 1-2-2a2 2 0 0 1 2 2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M27.945 38a6.553 6.553 0 0 0 6 0m21 22l-1.937-13.305l.937-1.25l6-2l5-6a1.414 1.414 0 0 0-2-2l-3 3c-1 1-4 0-7 1s-5 3.232-5 5V60M20.95 38c-3 0-4-7-4-14c0-6 5-12 14-12s14 6 14 12c0 7-1 14-4 14"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M19.886 31c-.102 7.999 4.936 14 11 14c5.936 0 11.063-6 11.063-14c0-5-3-11-3-11c-8 0-10 3-16 1c0 0-3 5-3.063 10Z"/>`),
		g.Group(children),
	)
}