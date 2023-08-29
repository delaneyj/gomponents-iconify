package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BirdOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M12.13 45.17a14.04 1.83 0 1 0 28.08 0a14.04 1.83 0 1 0-28.08 0Z" opacity=".15"/><path fill="#4aeff7" d="M23.366 39.194a18.31 17.93-76.8 1 0 8.362-35.652a18.31 17.93-76.8 1 0-8.362 35.652Z"/><path fill="#a6fbff" d="M10.1 22a17.87 17.87 0 0 1 35.24 1.6A18.08 18.08 0 0 0 31.73 3.55A18.11 18.11 0 0 0 10.1 17.28a18.84 18.84 0 0 0-.34 6.58c.09-.6.24-1.21.34-1.86Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M23.366 39.194a18.31 17.93-76.8 1 0 8.362-35.652a18.31 17.93-76.8 1 0-8.362 35.652Z"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M3.91 25.49s0 6.37 8 6.24c2.93-.19 4.47-7 4.47-7Z"/><path fill="#fff48c" d="M9.82 18.47s-1.87-.19-5 2.59A6.93 6.93 0 0 0 2.19 27s6.58 1.23 9.89.27s4.25-2.52 4.25-2.52s-.6-4.42-6.51-6.28Z"/><path fill="#fffacf" d="M7.06 22.87c3.75-2.94 6.13-2.56 6.13-2.56l.44.16a11.29 11.29 0 0 0-3.81-2s-1.87-.19-5 2.59A6.93 6.93 0 0 0 2.19 27s.82.15 2 .31a7.32 7.32 0 0 1 2.87-4.44Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M9.82 18.47s-1.87-.19-5 2.59A6.93 6.93 0 0 0 2.19 27s6.58 1.23 9.89.27s4.25-2.52 4.25-2.52s-.6-4.42-6.51-6.28Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24.37 10.59c-1.4-.29-4.4 0-6.15 4.83s-1.06 7.17 1.3 7.83s11.45-.88 12.23-4c.99-4.03-2.41-7.59-7.38-8.66Z"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M25.06 17.66a1.83 1.83 0 1 1-1.37-2.21a1.83 1.83 0 0 1 1.37 2.21Z"/><path fill="#d9fdff" d="M35.66 26.41c-.28.66-1.47 1.07-2.67.91s-1.93-.81-1.65-1.47s1.47-1.06 2.67-.91s1.93.82 1.65 1.47Z"/>`),
		g.Group(children),
	)
}