package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithSteamFromNose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M24 1.5a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37a18.25 18.25 0 1 1 18.25-18.25A18.25 18.25 0 0 1 24 38.5Z"/><path fill="#fff48c" d="M18 5.5a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#45413c" d="M8 45.5a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M10.75 15.5a11.43 11.43 0 0 1 3.37 1.1A11.35 11.35 0 0 1 17 18.66m20.25-3.16A11.65 11.65 0 0 0 31 18.66"/><path fill="#ffaa54" d="M38.5 26.5c0 .83-1.12 1.5-2.5 1.5s-2.5-.67-2.5-1.5S34.62 25 36 25s2.5.67 2.5 1.5Zm-29 0c0 .83 1.12 1.5 2.5 1.5s2.5-.67 2.5-1.5S13.38 25 12 25s-2.5.67-2.5 1.5Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" d="M19 31.5a9 9 0 0 1 10 0"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M38 20.5a1.75 1.75 0 0 1-3.5 0m-21 0a1.75 1.75 0 0 1-3.5 0"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M41.69 34.86a2.62 2.62 0 0 0-3.54-1.09a2.55 2.55 0 0 0-.85.74L31 26a1.17 1.17 0 0 0-2 1.08l3.55 10a2.55 2.55 0 0 0-1.09.3A2.62 2.62 0 0 0 33.82 42a4.05 4.05 0 0 0 6.88-3.65a2.62 2.62 0 0 0 .99-3.49Zm-35.38 0a2.62 2.62 0 0 1 3.54-1.09a2.55 2.55 0 0 1 .85.74L17 26a1.17 1.17 0 0 1 2 1.08L15.46 37a2.55 2.55 0 0 1 1.09.3a2.62 2.62 0 0 1-2.37 4.7a4.05 4.05 0 0 1-6.88-3.65a2.62 2.62 0 0 1-.99-3.49Z"/>`),
		g.Group(children),
	)
}