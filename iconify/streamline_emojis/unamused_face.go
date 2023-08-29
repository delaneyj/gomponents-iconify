package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnamusedFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M24 1.5a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37a18.25 18.25 0 1 1 18.25-18.25A18.25 18.25 0 0 1 24 38.5Z"/><path fill="#fff48c" d="M18 5.5a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="#45413c" d="M8 45.5a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Zm17.5 10h10"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M13.25 17.16A11.65 11.65 0 0 0 19.5 14m15.25 3.16A11.65 11.65 0 0 1 28.5 14"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M16 22.07a1.08 1.08 0 1 1-2.15 0C13.85 21.55 11 21 11 21h3.93A1.07 1.07 0 0 1 16 22.07Zm18 0a1.08 1.08 0 1 1-2.15 0C31.88 21.55 29 21 29 21h4a1.07 1.07 0 0 1 1 1.07Z"/><path fill="#ffaa54" d="M38 29c0 .83-1.12 1.5-2.5 1.5S33 29.83 33 29s1.12-1.5 2.5-1.5s2.5.67 2.5 1.5Zm-28 0c0 .83 1.12 1.5 2.5 1.5S15 29.83 15 29s-1.12-1.5-2.5-1.5s-2.5.67-2.5 1.5Z"/>`),
		g.Group(children),
	)
}