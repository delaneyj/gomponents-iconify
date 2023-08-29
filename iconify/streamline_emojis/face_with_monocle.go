package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithMonocle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M8 45.5a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="#ffe500" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M24 1.5a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37a18.25 18.25 0 1 1 18.25-18.25A18.25 18.25 0 0 1 24 38.5Z"/><path fill="#fff48c" d="M18 5.5a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ffaa54" d="M38.5 26.5c0 .83-1.12 1.5-2.5 1.5s-2.5-.67-2.5-1.5S34.62 25 36 25s2.5.67 2.5 1.5Zm-29 0c0 .83 1.12 1.5 2.5 1.5s2.5-.67 2.5-1.5S13.38 25 12 25s-2.5.67-2.5 1.5Z"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M14.5 17.5a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm19 3a1 1 0 1 0 1-1a1 1 0 0 0-1 1Z"/><path fill="#87898c" d="M42.5 14H44a1.5 1.5 0 0 1 1.5 1.5V17a1.5 1.5 0 0 1-1.5 1.5h-1.5V14Z"/><path fill="#bdbec0" d="M44 14h-1.5v1.5H44a1.5 1.5 0 0 1 1.5 1.5v-1.5A1.5 1.5 0 0 0 44 14Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M42.5 14H44a1.5 1.5 0 0 1 1.5 1.5V17a1.5 1.5 0 0 1-1.5 1.5h-1.5h0V14h0Z"/><path fill="#87898c" d="M26.5 16.5a8.5 8.5 0 1 0 17 0a8.5 8.5 0 1 0-17 0Z"/><path fill="#bdbec0" d="M35 10a8.5 8.5 0 0 1 8.44 7.5c0-.33.06-.66.06-1a8.5 8.5 0 0 0-17 0c0 .34 0 .67.06 1A8.5 8.5 0 0 1 35 10Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M26.5 16.5a8.5 8.5 0 1 0 17 0a8.5 8.5 0 1 0-17 0Z"/><path fill="#00b8f0" d="M28.5 16.5a6.5 6.5 0 1 0 13 0a6.5 6.5 0 1 0-13 0Z"/><path fill="#4acfff" d="M35 10a6.5 6.5 0 0 0-5.65 9.71l7.82-9.33A6.51 6.51 0 0 0 35 10Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M28.5 16.5a6.5 6.5 0 1 0 13 0a6.5 6.5 0 1 0-13 0Zm17-.5s1.5.06 1.5 2v17a3.72 3.72 0 0 1-4 4a3.68 3.68 0 0 1-4-4v-.27m-18-6.11a6.61 6.61 0 0 1 3-.62a6.61 6.61 0 0 1 3 .62M11 15s.42-2.5 2.5-2.5S16 15 16 15"/>`),
		g.Group(children),
	)
}