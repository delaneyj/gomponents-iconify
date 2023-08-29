package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellWithSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M10.26 45.21a13.74 1.79 0 1 0 27.48 0a13.74 1.79 0 1 0-27.48 0Z" opacity=".15"/><path fill="#fff48c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21.61 7.39a2.39 2.39 0 1 0 4.78 0a2.39 2.39 0 1 0-4.78 0Z"/><path fill="#ffe500" d="M40.73 33.67c0-3.08-3.39-5.13-4.78-7.76s-.75-7-2.39-11.95S27.39 8 24 8s-7.92 1-9.56 6s-1 9.31-2.39 12s-4.78 4.68-4.78 7.76Z"/><path fill="#fff48c" d="M24 12.17c3.68 0 8.58 1 10.48 5.72a25.15 25.15 0 0 0-.92-3.89C31.92 9 27.39 8 24 8s-7.92 1-9.56 6a25.15 25.15 0 0 0-.92 3.93c1.9-4.79 6.8-5.76 10.48-5.76Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M40.73 33.67c0-3.08-3.39-5.13-4.78-7.76s-.75-7-2.39-11.95S27.39 8 24 8s-7.92 1-9.56 6s-1 9.31-2.39 12s-4.78 4.68-4.78 7.76Z"/><path fill="#ffe500" d="M7.27 33.67a16.73 1.79 0 1 0 33.46 0a16.73 1.79 0 1 0-33.46 0Z"/><path fill="#ebcb00" d="M9.51 34.57a92 92 0 0 1 14.49-.9a92 92 0 0 1 14.49.9c1.42-.27 2.24-.57 2.24-.9c0-1-7.49-1.79-16.73-1.79s-16.73.8-16.73 1.79c0 .33.82.63 2.24.9Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M7.27 33.67a16.73 1.79 0 1 0 33.46 0a16.73 1.79 0 1 0-33.46 0Z"/><path fill="#ffe500" d="M24 31.88h-3.11a3.58 3.58 0 1 0 6.22 0c-1.01.01-2.04 0-3.11 0Z"/><path fill="#ebcb00" d="M24 34.17h3.53a3.17 3.17 0 0 0 0-.53a3.57 3.57 0 0 0-.47-1.75h-6.22a3.57 3.57 0 0 0-.47 1.75a3.17 3.17 0 0 0 0 .53c1.25.01 2.42 0 3.63 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 31.88h-3.11a3.58 3.58 0 1 0 6.22 0c-1.01.01-2.04 0-3.11 0Z"/><path fill="#ff6242" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M43.28 38.89L8.22 3.83A1.5 1.5 0 1 0 6.1 5.94L41.16 41a1.5 1.5 0 0 0 2.12-2.11Z"/>`),
		g.Group(children),
	)
}