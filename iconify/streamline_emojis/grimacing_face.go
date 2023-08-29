package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrimacingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M24 1.5a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37a18.25 18.25 0 1 1 18.25-18.25A18.25 18.25 0 0 1 24 38.5Z"/><path fill="#fff48c" d="M18 5.5a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="#45413c" d="M8 45.5a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ffaa54" d="M38.5 26.5c0 .83-1.12 1.5-2.5 1.5s-2.5-.67-2.5-1.5S34.62 25 36 25s2.5.67 2.5 1.5Zm-29 0c0 .83 1.12 1.5 2.5 1.5s2.5-.67 2.5-1.5S13.38 25 12 25s-2.5.67-2.5 1.5Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 36.55c9.2 0 10.43-3.84 9.89-7a3.41 3.41 0 0 0-4.56-2.36a16.42 16.42 0 0 1-5.33.84a16.42 16.42 0 0 1-5.33-.84a3.41 3.41 0 0 0-4.56 2.36c-.54 3.16.69 7 9.89 7Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M34 30.28a17.08 17.08 0 0 1-4.6 1.36a32.3 32.3 0 0 1-5.38.41a32.3 32.3 0 0 1-5.38-.41A17.08 17.08 0 0 1 14 30.28"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m19.27 27.39l-.65 4.25l-.63 4.09m6.01.82v-8.5m4.73-.66l.65 4.25l.63 4.09"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M14.5 20.5a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm19 0a1 1 0 1 0 1-1a1 1 0 0 0-1 1Z"/>`),
		g.Group(children),
	)
}