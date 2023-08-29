package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrinningFaceWithSweat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M24 1.5a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37a18.25 18.25 0 1 1 18.25-18.25A18.25 18.25 0 0 1 24 38.5Z"/><path fill="#fff48c" d="M18 5.5a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="#45413c" d="M8 45.5a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ffaa54" d="M38.5 26.5c0 .83-1.12 1.5-2.5 1.5s-2.5-.67-2.5-1.5S34.62 25 36 25s2.5.67 2.5 1.5Zm-29 0c0 .83 1.12 1.5 2.5 1.5s2.5-.67 2.5-1.5S13.38 25 12 25s-2.5.67-2.5 1.5Z"/><path fill="#ffb0ca" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M17.94 29.5a.94.94 0 0 0-.71.33a1 1 0 0 0-.22.76a7.09 7.09 0 0 0 14 0a1 1 0 0 0-.22-.76a.94.94 0 0 0-.71-.33Z"/><path fill="#ff87af" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M29.57 33.84A11.37 11.37 0 0 0 24 32.53a11.37 11.37 0 0 0-5.57 1.31a7.16 7.16 0 0 0 11.14 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M10 22.75a1.75 1.75 0 0 1 3.5 0m21 0a1.75 1.75 0 0 1 3.5 0"/><path fill="#00b8f0" d="M47.5 33.5a5 5 0 0 1-10 0c0-2.76 5-11 5-11s5 8.24 5 11Z"/><path fill="#009fd9" d="M42.5 36a5 5 0 0 1-4.8-3.69a4.53 4.53 0 0 0-.2 1.19a5 5 0 0 0 10 0a4.53 4.53 0 0 0-.2-1.19A5 5 0 0 1 42.5 36Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M47.5 33.5a5 5 0 0 1-10 0c0-2.76 5-11 5-11s5 8.24 5 11Z"/>`),
		g.Group(children),
	)
}