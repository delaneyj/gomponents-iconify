package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CryingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M24 1.5a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37a18.25 18.25 0 1 1 18.25-18.25A18.25 18.25 0 0 1 24 38.5Z"/><path fill="#fff48c" d="M18 5.5a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" d="M19 31.5a9 9 0 0 1 10 0"/><path fill="#45413c" d="M8 45.5a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M15 21.5a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm18 0a1 1 0 1 0 1-1a1 1 0 0 0-1 1Z"/><path fill="#ffaa54" d="M38.5 26.5c0 .83-1.12 1.5-2.5 1.5s-2.5-.67-2.5-1.5S34.62 25 36 25s2.5.67 2.5 1.5Zm-29 0c0 .83 1.12 1.5 2.5 1.5s2.5-.67 2.5-1.5S13.38 25 12 25s-2.5.67-2.5 1.5Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M8.7 20.93a7 7 0 0 1 5.39-5.23a6.92 6.92 0 0 1 3 0m22.21 5.23a7 7 0 0 0-5.39-5.23a6.92 6.92 0 0 0-3 0"/><path fill="#00b8f0" d="M37.89 40.28c-3.65-1.6.85-5.6.35-8.35c4.14 2.58 3.29 9.94-.35 8.35Z"/><path fill="#4acfff" d="M38.24 31.93a1.2 1.2 0 0 1 0 .2c2.68 3.06 1.7 8.62-1.46 7.36a2.61 2.61 0 0 0 1.08.79c3.67 1.59 4.52-5.77.38-8.35Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M37.89 40.28c-3.65-1.6.85-5.6.35-8.35c4.14 2.58 3.29 9.94-.35 8.35Z"/><path fill="#00b8f0" d="M37.71 29.81c-2.75-.58-.22-4-1-5.81c3.29 1.19 3.74 6.39 1 5.81Z"/><path fill="#4acfff" d="M36.76 24v.14c2.29 1.72 2.41 5.7.05 5.28a1.81 1.81 0 0 0 .86.39c2.78.58 2.33-4.62-.91-5.81Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M37.71 29.81c-2.75-.58-.22-4-1-5.81c3.29 1.19 3.74 6.39 1 5.81Z"/><path fill="#00b8f0" d="M10.29 29.81c2.75-.58.22-4 1-5.81c-3.29 1.19-3.74 6.39-1 5.81Z"/><path fill="#4acfff" d="M11.24 24v.14c-2.29 1.72-2.41 5.7 0 5.28a1.81 1.81 0 0 1-.86.39c-2.83.58-2.38-4.62.86-5.81Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M10.29 29.81c2.75-.58.22-4 1-5.81c-3.29 1.19-3.74 6.39-1 5.81Z"/>`),
		g.Group(children),
	)
}