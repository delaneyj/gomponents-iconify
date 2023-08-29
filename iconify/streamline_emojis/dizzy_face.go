package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DizzyFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M24 1.5a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37a18.25 18.25 0 1 1 18.25-18.25A18.25 18.25 0 0 1 24 38.5Z"/><path fill="#fff48c" d="M18 5.5a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="#45413c" d="M8 45.5a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ffaa54" d="M38.5 26.5c0 .83-1.12 1.5-2.5 1.5s-2.5-.67-2.5-1.5S34.62 25 36 25s2.5.67 2.5 1.5Zm-29 0c0 .83 1.12 1.5 2.5 1.5s2.5-.67 2.5-1.5S13.38 25 12 25s-2.5.67-2.5 1.5Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m12 19l3.5 3.5m0-3.5L12 22.5M36 19l-3.5 3.5m0-3.5l3.5 3.5M10.67 16a4.46 4.46 0 0 1 2.83-1a4.51 4.51 0 0 1 2.83 1m21 0a4.46 4.46 0 0 0-2.83-1a4.51 4.51 0 0 0-2.83 1"/><path fill="#ffb0ca" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M16.5 31.5a7.5 7.5 0 1 0 15 0a7.5 7.5 0 1 0-15 0Z"/><path fill="#ff866e" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M17.65 35.48a6.35 3.52 0 1 0 12.7 0a6.35 3.52 0 1 0-12.7 0Z"/>`),
		g.Group(children),
	)
}