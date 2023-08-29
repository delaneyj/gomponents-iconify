package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectHit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ff6242" d="M5.39 24a18.61 18.61 0 1 0 37.22 0a18.61 18.61 0 1 0-37.22 0Z"/><path fill="#ff866e" d="M24 9.18A18.62 18.62 0 0 1 42.52 25.9c.06-.63.09-1.26.09-1.9a18.61 18.61 0 0 0-37.22 0c0 .64 0 1.27.09 1.9A18.62 18.62 0 0 1 24 9.18Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M5.39 24a18.61 18.61 0 1 0 37.22 0a18.61 18.61 0 1 0-37.22 0Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M9.9 24a14.1 14.1 0 1 0 28.2 0a14.1 14.1 0 1 0-28.2 0Z"/><path fill="#ff6242" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M14.41 24a9.59 9.59 0 1 0 19.18 0a9.59 9.59 0 1 0-19.18 0Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M19.49 24a4.51 4.51 0 1 0 9.02 0a4.51 4.51 0 1 0-9.02 0Z"/><path fill="#00b8f0" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M38.24 38a1.12 1.12 0 0 1-1.59 0L23.49 24.8a1.13 1.13 0 0 1 0-1.6a1.12 1.12 0 0 1 1.59 0l13.16 13.16a1.13 1.13 0 0 1 0 1.64Z"/><path fill="#00b8f0" d="M30.27 28.38s6.65-1.76 10.37.8c3.54 2.45 1.72 5.76-.8 6.38c-2.08.52-4-1.59-4-1.59Z"/><path fill="#4acfff" d="M40.64 32a5 5 0 0 1 1.61 1.73c.69-1.32.49-3.07-1.61-4.52c-3.72-2.56-10.37-.8-10.37-.8l2.36 2.36c2.27-.31 5.71-.39 8.01 1.23Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M30.27 28.38s6.65-1.76 10.37.8c3.54 2.45 1.72 5.76-.8 6.38c-2.08.52-4-1.59-4-1.59Z"/><path fill="#00b8f0" d="M28.68 30s-1.76 6.65.8 10.37c2.44 3.54 5.76 1.72 6.38-.8c.52-2.09-1.59-4-1.59-4Z"/><path fill="#009fd9" d="M35.53 40.44a3.81 3.81 0 0 0 .33-.87c.52-2.09-1.59-4-1.59-4L28.68 30a21.92 21.92 0 0 0-.51 3.09Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M28.68 30s-1.76 6.65.8 10.37c2.44 3.54 5.76 1.72 6.38-.8c.52-2.09-1.59-4-1.59-4Z"/><path fill="#45413c" d="M11.03 45.31a12.97 1.69 0 1 0 25.94 0a12.97 1.69 0 1 0-25.94 0Z" opacity=".15"/>`),
		g.Group(children),
	)
}