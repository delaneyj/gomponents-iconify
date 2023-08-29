package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rainbow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M3.08 45.3a20.92 1.7 0 1 0 41.84 0a20.92 1.7 0 1 0-41.84 0Z" opacity=".15"/><path fill="#fff" d="M4.21 4.21h39.57v39.57H4.21Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4.21 4.21h39.57v39.57H4.21Z"/><path fill="#00b8f0" d="M7.04 7.04h33.92v33.92H7.04Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M7.04 7.04h33.92v33.92H7.04Z"/><path fill="#ff6242" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M41 9.4h0v30.43A1.13 1.13 0 0 1 39.83 41H9.4h0v0A31.56 31.56 0 0 1 41 9.4Z"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M41 14.13h0v25.7A1.13 1.13 0 0 1 39.83 41h-25.7h0v0A26.83 26.83 0 0 1 41 14.13Z"/><path fill="#00dba8" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M41 19h0v20.83A1.13 1.13 0 0 1 39.83 41H19h0v0a22 22 0 0 1 22-22Z"/><path fill="#bf8df2" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M41 23.73h0v16.1A1.13 1.13 0 0 1 39.83 41h-16.1h0v0A17.23 17.23 0 0 1 41 23.73Z"/><path fill="#00b8f0" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M41 28.25h0v11.58A1.13 1.13 0 0 1 39.83 41H28.25h0v0A12.71 12.71 0 0 1 41 28.25Z"/>`),
		g.Group(children),
	)
}