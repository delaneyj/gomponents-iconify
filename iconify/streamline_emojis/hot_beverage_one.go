package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotBeverageOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M9.5 44.39a14.5 1.61 0 1 0 29 0a14.5 1.61 0 1 0-29 0Z" opacity=".15"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M15.49 10.08a4.72 4.72 0 1 0-4.08 2.34h.44v0a2.58 2.58 0 1 0 3.65-2.34Zm4.01-7a2.58 2.58 0 1 0 5.16 0a2.58 2.58 0 1 0-5.16 0Z"/><path fill="#e5f8ff" d="M40.72 34.22a5.59 5.59 0 1 0 0-11.17h-2.15v3.44h2.15a2.15 2.15 0 0 1 0 4.3h-3.86v3.43Z"/><path fill="#fff" d="M39.86 24.77a5.59 5.59 0 0 1 5.59 5.59a6 6 0 0 1-.23 1.64a5.59 5.59 0 0 0-4.5-8.9h-2.15v1.72Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M40.72 34.22a5.59 5.59 0 1 0 0-11.17h-2.15v3.44h2.15a2.15 2.15 0 0 1 0 4.3h-3.86v3.43Z"/><path fill="#fff" d="M8.5 19.5v9a15.75 15.75 0 0 0 15.5 16a15.75 15.75 0 0 0 15.5-16v-9Z"/><path fill="#e8f4fa" d="M24 39c-6.45 0-12.08-2.79-15.11-6.94A15.62 15.62 0 0 0 24 44.5a15.62 15.62 0 0 0 15.11-12.44C36.08 36.21 30.45 39 24 39Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M8.5 19.5v9a15.75 15.75 0 0 0 15.5 16a15.75 15.75 0 0 0 15.5-16v-9Z"/><path fill="#e8f4fa" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M8.5 19.5a15.5 2 0 1 0 31 0a15.5 2 0 1 0-31 0Z"/><path fill="#ffb0ca" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M27.32 32.94a.86.86 0 0 1 .84 1.06a4.3 4.3 0 0 1-8.32 0a.86.86 0 0 1 .84-1.07Z"/><path fill="#ff87af" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 37.23a4.26 4.26 0 0 0 3.38-1.67a6.26 6.26 0 0 0-3.38-.91a6.26 6.26 0 0 0-3.38.91A4.26 4.26 0 0 0 24 37.23Z"/><path fill="#fcd" d="M15.41 30.79c0 .47.57.86 1.29.86s1.3-.39 1.3-.86s-.58-.86-1.29-.86s-1.3.38-1.3.86Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M16.27 29.07a1.29 1.29 0 1 1 2.57 0"/><path fill="#fcd" d="M32.59 30.79c0 .47-.57.86-1.29.86s-1.3-.39-1.3-.86s.58-.86 1.29-.86s1.3.38 1.3.86Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M31.73 29.07a1.29 1.29 0 1 0-2.57 0"/>`),
		g.Group(children),
	)
}