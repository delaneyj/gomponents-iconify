package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomanPoutingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M9 45.5a15 1.5 0 1 0 30 0a15 1.5 0 1 0-30 0Z" opacity=".15"/><path fill="#724f3d" d="M34.88 31.38a1.24 1.24 0 0 1-.77 1.14a27.09 27.09 0 0 1-20.22 0a1.24 1.24 0 0 1-.77-1.14V17.19a10.88 10.88 0 0 1 21.76 0Z"/><path fill="#a86c4d" d="M24 6.32a10.88 10.88 0 0 0-10.88 10.87v2.9a10.88 10.88 0 0 1 21.76 0v-2.9A10.88 10.88 0 0 0 24 6.32Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M34.88 31.38a1.24 1.24 0 0 1-.77 1.14h0a27.09 27.09 0 0 1-20.22 0h0a1.24 1.24 0 0 1-.77-1.14V17.19a10.88 10.88 0 0 1 21.76 0Z"/><path fill="#ff87af" d="M35.88 45H12.12v-1.9a11.88 11.88 0 1 1 23.76 0Z"/><path fill="#ff6196" d="M24 31.22A11.88 11.88 0 0 0 12.12 43.1V45h.17a11.88 11.88 0 0 1 23.42 0h.17v-1.9A11.88 11.88 0 0 0 24 31.22Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M35.88 45H12.12v-1.9a11.88 11.88 0 1 1 23.76 0Z"/><path fill="#ffe500" d="M24 27.17a2.54 2.54 0 0 0-2.54 2.55v2.85a2.54 2.54 0 1 0 5.08 0v-2.85A2.54 2.54 0 0 0 24 27.17Z"/><path fill="#ebcb00" d="M24 27.17a2.54 2.54 0 0 0-2.54 2.54v.6a2.54 2.54 0 0 0 5.08 0v-.59A2.54 2.54 0 0 0 24 27.17Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 27.17a2.54 2.54 0 0 0-2.54 2.55v2.85a2.54 2.54 0 1 0 5.08 0v-2.85A2.54 2.54 0 0 0 24 27.17Z"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33.06 19.47a8.77 8.77 0 0 1-5.67-5L26.77 13a15.92 15.92 0 0 1-9.95 6.18l-1.89.34a1.81 1.81 0 0 0 0 3.62h.12a9 9 0 0 0 17.9 0h.11a1.81 1.81 0 1 0 0-3.62Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M17.06 45v-2.41M30.94 45v-2.41"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M18.57 20.91a.77.77 0 1 0 .77-.77a.76.76 0 0 0-.77.77Zm10.86 0a.77.77 0 1 1-.77-.77a.76.76 0 0 1 .77.77Z"/><path fill="#ffaa54" d="M17.06 24.19a1 .6 0 1 0 2 0a1 .6 0 1 0-2 0Zm11.88 0a1 .6 0 1 0 2 0a1 .6 0 1 0-2 0Z"/><path fill="#ff866e" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M23.06 28.54a.5.5 0 0 1-.5-.53c0-.79.29-2.21 1.44-2.21s1.4 1.42 1.44 2.21a.5.5 0 0 1-.5.53Z"/>`),
		g.Group(children),
	)
}