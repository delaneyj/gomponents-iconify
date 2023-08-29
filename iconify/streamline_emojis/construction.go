package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Construction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M2.42 42.98a21.58 2.02 0 1 0 43.16 0a21.58 2.02 0 1 0-43.16 0Z" opacity=".15"/><path fill="#ffe500" d="M2.25 7.93h43.5v17.53H2.25Z"/><path fill="#fff48c" d="M44.4 7.93H3.6a1.34 1.34 0 0 0-1.35 1.35v3.46a1.34 1.34 0 0 1 1.35-1.35h40.8a1.34 1.34 0 0 1 1.35 1.35V9.28a1.34 1.34 0 0 0-1.35-1.35Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M2.25 7.93h43.5v17.53H2.25Z"/><path fill="#daedf7" d="M6.72 25.47h5.4v16.18A1.35 1.35 0 0 1 10.76 43H8.07a1.35 1.35 0 0 1-1.35-1.35V25.47Zm29.17 0h5.4v16.18A1.35 1.35 0 0 1 39.93 43h-2.7a1.35 1.35 0 0 1-1.35-1.35V25.47h.01Z"/><path fill="#c0dceb" d="M6.72 25.47h5.4v3.26h-5.4zm29.17 0h5.4v3.26h-5.4z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M6.72 25.47h5.4v16.18A1.35 1.35 0 0 1 10.76 43H8.07a1.35 1.35 0 0 1-1.35-1.35V25.47h0Zm29.17 0h5.4v16.18A1.35 1.35 0 0 1 39.93 43h-2.7a1.35 1.35 0 0 1-1.35-1.35V25.47h.01Z"/><path fill="#656769" d="M8.84 7.93L3.2 25.41a1.76 1.76 0 0 0 .4.06h5.57l5.65-17.54Zm18.12 0h-5.98l-5.66 17.53h5.99l5.65-17.53zm12.14 0h-5.98l-5.66 17.53h5.99L39.1 7.93zm6.08.25L39.6 25.47h4.8a1.35 1.35 0 0 0 1.35-1.35V9.28a1.34 1.34 0 0 0-.57-1.1Z"/><path fill="#87898c" d="m13.71 11.38l1.11-3.45H8.84l-1.12 3.45h5.99zm12.14 0l1.11-3.45h-5.98l-1.12 3.45h5.99zm12.14 0l1.11-3.45h-5.98L32 11.38h5.99zm7.19-3.2l-1 3.2h.26a1.35 1.35 0 0 1 1.35 1.35V9.28a1.34 1.34 0 0 0-.61-1.1Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M8.84 7.93L3.2 25.41a1.76 1.76 0 0 0 .4.06h5.57l5.65-17.54Zm18.12 0h-5.98l-5.66 17.53h5.99l5.65-17.53zm12.14 0h-5.98l-5.66 17.53h5.99L39.1 7.93zm6.08.25L39.6 25.47h4.8a1.35 1.35 0 0 0 1.35-1.35V9.28a1.34 1.34 0 0 0-.57-1.1Z"/>`),
		g.Group(children),
	)
}