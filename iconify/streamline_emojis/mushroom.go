package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mushroom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M11.61 44.88a12.4 1.62 0 1 0 24.8 0a12.4 1.62 0 1 0-24.8 0Z" opacity=".15"/><path fill="#fff5e3" d="M31 33.68a50.23 50.23 0 0 1-.54-7.55H17.53a50.23 50.23 0 0 1-.53 7.55c-.62 4.58-2.78 6.2-2.69 8.62S18.79 45 24 45s9.61-.27 9.7-2.7s-2.07-4.04-2.7-8.62Z"/><path fill="#f7e5c6" d="M16.77 34.94a24.1 24.1 0 0 0 14.46 0c-.09-.4-.16-.82-.22-1.26a50.23 50.23 0 0 1-.54-7.55H17.53a50.23 50.23 0 0 1-.53 7.55c-.07.44-.14.86-.23 1.26Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M31 33.68a50.23 50.23 0 0 1-.54-7.55H17.53a50.23 50.23 0 0 1-.53 7.55c-.62 4.58-2.78 6.2-2.69 8.62S18.79 45 24 45s9.61-.27 9.7-2.7s-2.07-4.04-2.7-8.62Z"/><path fill="#ff6242" d="M24 3.48c-13.18 0-21.95 12.73-20.95 19.58s11 7.38 21 7.38s19.94-.54 21-7.38S37.18 3.48 24 3.48Z"/><path fill="#ff866e" d="M3.6 24.9c2.18-6.76 9.89-14.85 20.4-14.85s18.22 8.09 20.4 14.85a6.68 6.68 0 0 0 .6-1.84c1-6.85-7.77-19.58-21-19.58S2.05 16.21 3.05 23.06a6.68 6.68 0 0 0 .55 1.84Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 3.48c-13.18 0-21.95 12.73-20.95 19.58s11 7.38 21 7.38s19.94-.54 21-7.38S37.18 3.48 24 3.48Z"/><path fill="#fff48c" d="M32 5.11c-3.36-.11-6 1.17-6.36 3.38c-.48 2.77 2.77 5.92 7.24 7c4.18 1 7.95-.06 8.82-2.46A23.29 23.29 0 0 0 32 5.11Z"/><path fill="#fffacf" d="M38.18 15.67a4.6 4.6 0 0 0 3.52-2.61A23.29 23.29 0 0 0 32 5.11c-3.36-.11-6 1.17-6.36 3.38a3.54 3.54 0 0 0 .12 1.64a21.38 21.38 0 0 1 12.42 5.54Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M32 5.11c-3.36-.11-6 1.17-6.36 3.38c-.48 2.77 2.77 5.92 7.24 7c4.18 1 7.95-.06 8.82-2.46A23.29 23.29 0 0 0 32 5.11Z"/><path fill="#fff48c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M9 15c-2 1.35-3 3.89-2.14 5.72S10 23 12 21.65s2.93-3.86 2.1-5.69S10.94 13.72 9 15Zm20.82 8.43a3.14 3.14 0 1 1-3.14-3.24a3.19 3.19 0 0 1 3.14 3.24Z"/><path fill="#fffacf" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M18.62 8.64A1.79 1.79 0 0 1 17 10.48a1.64 1.64 0 0 1-1.94-1.38a1.81 1.81 0 0 1 1.64-1.84a1.64 1.64 0 0 1 1.92 1.38Z"/><path fill="#fff48c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M38.32 22.35a1.79 1.79 0 1 0 1.78-1.62a1.7 1.7 0 0 0-1.78 1.62Z"/>`),
		g.Group(children),
	)
}