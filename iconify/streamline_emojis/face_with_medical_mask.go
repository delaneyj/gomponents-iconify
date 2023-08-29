package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithMedicalMask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M4 21.25a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M24 1.25a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37A18.25 18.25 0 1 1 42.25 20A18.25 18.25 0 0 1 24 38.25Z"/><path fill="#fff48c" d="M18 5.25a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="#45413c" d="M8 45.25a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="#ffaa54" d="M38.5 26.25c0 .83-1.12 1.5-2.5 1.5s-2.5-.67-2.5-1.5s1.12-1.5 2.5-1.5s2.5.67 2.5 1.5Zm-29 0c0 .83 1.12 1.5 2.5 1.5s2.5-.67 2.5-1.5s-1.12-1.5-2.5-1.5s-2.5.67-2.5 1.5Z"/><path fill="#8cffe4" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M13 38a20 20 0 0 0 22 0v-8.7s-6-2.5-11-2.5s-11 2.5-11 2.5Z"/><path fill="#e5fff9" d="M24 26.75c-5 0-11 2.5-11 2.5v2.5s6-2.5 11-2.5s11 2.5 11 2.5v-2.5s-6-2.5-11-2.5Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M13 38a20 20 0 0 0 22 0v-8.7s-6-2.5-11-2.5s-11 2.5-11 2.5Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 21.25a20 20 0 1 0 40 0a20 20 0 1 0-40 0Zm9 8l-9-7m31 7l9-7"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M35 33.25a7.69 7.69 0 0 0 4 1m-26-1a7.69 7.69 0 0 1-4 1m7.5-13.5H11m26 0h-5.5M13 15.91a11.43 11.43 0 0 0 3.37-1.1a11.18 11.18 0 0 0 2.88-2.06M35 15.91a11.43 11.43 0 0 1-3.37-1.1a11.18 11.18 0 0 1-2.88-2.06"/>`),
		g.Group(children),
	)
}