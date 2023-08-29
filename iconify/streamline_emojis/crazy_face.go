package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrazyFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M24 1.5a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37a18.25 18.25 0 1 1 18.25-18.25A18.25 18.25 0 0 1 24 38.5Z"/><path fill="#fff48c" d="M18 5.5a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="#45413c" d="M8 45.5a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ffaa54" d="M38.5 26.5c0 .83-1.12 1.5-2.5 1.5s-2.5-.67-2.5-1.5S34.62 25 36 25s2.5.67 2.5 1.5Zm-29 0c0 .83 1.12 1.5 2.5 1.5s2.5-.67 2.5-1.5S13.38 25 12 25s-2.5.67-2.5 1.5Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 38.5a10.9 10.9 0 0 0 11-9.17a1 1 0 0 0-.31-.84a1.12 1.12 0 0 0-.91-.27a61.58 61.58 0 0 1-19.54 0a1.12 1.12 0 0 0-.91.27a1 1 0 0 0-.31.84A10.9 10.9 0 0 0 24 38.5Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M34 32.5A30.23 30.23 0 0 1 24 34a30.23 30.23 0 0 1-10-1.5"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 29v9.5m-6.5-9.84c0 .6-.05 5.75 1.06 8.49m11.94-8.49c0 .6.05 5.75-1.06 8.49"/><path fill="#fff" d="M28.5 14.5a8 8 0 1 0 16 0a8 8 0 1 0-16 0Zm-21 3a5 5 0 1 0 10 0a5 5 0 1 0-10 0Z"/><path fill="#f0f0f0" d="M36.5 19.5a8 8 0 0 1-7.85-6.5a7.48 7.48 0 0 0-.15 1.5a8 8 0 0 0 16 0a7.48 7.48 0 0 0-.15-1.5a8 8 0 0 1-7.85 6.5Zm-24 1.12a5 5 0 0 1-4.91-4.06a5.87 5.87 0 0 0-.09.94a5 5 0 0 0 10 0a5.87 5.87 0 0 0-.09-.94a5 5 0 0 1-4.91 4.06Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M28.5 14.5a8 8 0 1 0 16 0a8 8 0 1 0-16 0Zm-21 3a5 5 0 1 0 10 0a5 5 0 1 0-10 0Z"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M13.5 17.5a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm22-3a1 1 0 1 0 1-1a1 1 0 0 0-1 1Z"/>`),
		g.Group(children),
	)
}