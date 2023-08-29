package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WinkingFaceWithTongue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M24 1.5a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37a18.25 18.25 0 1 1 18.25-18.25A18.25 18.25 0 0 1 24 38.5Z"/><path fill="#fff48c" d="M18 5.5a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#45413c" d="M8 45.5a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="#ffaa54" d="M40 26.5c0 .83-1.12 1.5-2.5 1.5s-2.5-.67-2.5-1.5s1.12-1.5 2.5-1.5s2.5.67 2.5 1.5Zm-32 0c0 .83 1.12 1.5 2.5 1.5s2.5-.67 2.5-1.5s-1.12-1.5-2.5-1.5s-2.5.67-2.5 1.5Z"/><path fill="#ffb0ca" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M16.47 27.93c.5 8 5.09 11 6.85 12a1.37 1.37 0 0 0 1.36 0c1.76-1 6.35-3.92 6.85-12a1.37 1.37 0 0 0-1.38-1.45h-12.3a1.37 1.37 0 0 0-1.38 1.45Z"/><path fill="#ff6242" d="M24 33c0-1.5 4.5-1.06 4.5.5V40a4.5 4.5 0 0 1-9 0v-6.5c0-1.56 4.5-2 4.5-.5Z"/><path fill="#ff866e" d="M24 33c0-1.5-4.5-1.06-4.5.5v2c0-1.56 9-1.56 9 0v-2c0-1.56-4.5-2-4.5-.5Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 33c0-1.5 4.5-1.06 4.5.5V40a4.5 4.5 0 0 1-9 0v-6.5c0-1.56 4.5-2 4.5-.5Zm12.26-15l-3.21 1.61a1 1 0 0 0 0 1.78L36.26 23m-24.52-5L15 19.61a1 1 0 0 1 0 1.78L11.74 23"/>`),
		g.Group(children),
	)
}