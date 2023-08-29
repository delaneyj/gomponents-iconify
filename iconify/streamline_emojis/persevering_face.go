package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PerseveringFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M24 1.5a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37a18.25 18.25 0 1 1 18.25-18.25A18.25 18.25 0 0 1 24 38.5Z"/><path fill="#fff48c" d="M18 5.5a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="#45413c" d="M8 44.91a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ffaa54" d="M38.5 26.5c0 .83-1.12 1.5-2.5 1.5s-2.5-.67-2.5-1.5S34.62 25 36 25s2.5.67 2.5 1.5Zm-29 0c0 .83 1.12 1.5 2.5 1.5s2.5-.67 2.5-1.5S13.38 25 12 25s-2.5.67-2.5 1.5Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m36.26 18l-3.21 1.61a1 1 0 0 0 0 1.78L36.26 23m-24.52-5L15 19.61a1 1 0 0 1 0 1.78L11.74 23M13 16a11.43 11.43 0 0 0 3.37-1.1a11.18 11.18 0 0 0 2.88-2.06M35 16a11.43 11.43 0 0 1-3.37-1.1a11.18 11.18 0 0 1-2.88-2.06"/><path fill="none" stroke="#45413c" stroke-linecap="round" d="M19 31.5a9 9 0 0 1 10 0"/>`),
		g.Group(children),
	)
}