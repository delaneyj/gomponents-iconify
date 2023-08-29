package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShushingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M24 1.5a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37a18.25 18.25 0 1 1 18.25-18.25A18.25 18.25 0 0 1 24 38.5Z"/><path fill="#fff48c" d="M18 5.5a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="#45413c" d="M8 45.5a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ffaa54" d="M38.5 24.5c0 .83-1.12 1.5-2.5 1.5s-2.5-.67-2.5-1.5S34.62 23 36 23s2.5.67 2.5 1.5Zm-29 0c0 .83 1.12 1.5 2.5 1.5s2.5-.67 2.5-1.5S13.38 23 12 23s-2.5.67-2.5 1.5Z"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M14.5 18.5a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm19 0a1 1 0 1 0 1-1a1 1 0 0 0-1 1Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M18.5 27a19 19 0 0 0 5.5 1a19 19 0 0 0 5.5-1"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 22.5a2 2 0 0 0-2 2V36h4V24.5a2 2 0 0 0-2-2Z"/><path fill="#ffe500" d="M20.4 45.5a6 6 0 0 0 3.73-1.32l5.1-4.18a2 2 0 0 0-1.88-3.52l-1.35.44v-3A3.89 3.89 0 0 0 22.11 30H20a2.5 2.5 0 0 0-2.5 2.5a3 3 0 0 0-3 3v4.1a5.9 5.9 0 0 0 5.9 5.9Z"/><path fill="#ebcb00" d="m29.23 38l-5.1 4.15a6 6 0 0 1-3.73 1.35a5.9 5.9 0 0 1-5.9-5.9v2a5.9 5.9 0 0 0 5.9 5.9a6 6 0 0 0 3.73-1.32l5.1-4.18a2 2 0 0 0 .54-2.49s-.33.35-.54.49Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M20.4 45.5a6 6 0 0 0 3.73-1.32l5.1-4.18a2 2 0 0 0-1.88-3.52l-1.35.44v-3A3.89 3.89 0 0 0 22.11 30H20a2.5 2.5 0 0 0-2.5 2.5a3 3 0 0 0-3 3v4.1a5.9 5.9 0 0 0 5.9 5.9Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M18.5 36v3.5a2 2 0 0 0 4 0v-1.75"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M26 33.89v3.86a1.75 1.75 0 0 1-3.5 0v-2.22a3 3 0 0 0-3-3h-2"/>`),
		g.Group(children),
	)
}