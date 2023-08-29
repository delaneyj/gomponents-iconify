package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SadButRelievedFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M24 1.5a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37a18.25 18.25 0 1 1 18.25-18.25A18.25 18.25 0 0 1 24 38.5Z"/><path fill="#fff48c" d="M18 5.5a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" d="M19 31.5a9 9 0 0 1 10 0"/><path fill="#45413c" d="M8 45.5a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M15 21.5a1 1 0 1 1-1-1a1 1 0 0 1 1 1Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M13 17.16a11.43 11.43 0 0 0 3.37-1.1A11.18 11.18 0 0 0 19.25 14"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33 21.5a1 1 0 1 0 1-1a1 1 0 0 0-1 1Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M35 17.16a11.43 11.43 0 0 1-3.37-1.1A11.18 11.18 0 0 1 28.75 14"/><path fill="#ffaa54" d="M38.5 26.5c0 .83-1.12 1.5-2.5 1.5s-2.5-.67-2.5-1.5S34.62 25 36 25s2.5.67 2.5 1.5Zm-29 0c0 .83 1.12 1.5 2.5 1.5s2.5-.67 2.5-1.5S13.38 25 12 25s-2.5.67-2.5 1.5Z"/><path fill="#00b8f0" d="M46.5 29a5 5 0 0 1-10 0c0-2.76 5-11 5-11s5 8.24 5 11Z"/><path fill="#009fd9" d="M41.5 31.5a5 5 0 0 1-4.8-3.69a4.53 4.53 0 0 0-.2 1.19a5 5 0 0 0 10 0a4.53 4.53 0 0 0-.2-1.19a5 5 0 0 1-4.8 3.69Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M46.5 29a5 5 0 0 1-10 0c0-2.76 5-11 5-11s5 8.24 5 11Z"/>`),
		g.Group(children),
	)
}