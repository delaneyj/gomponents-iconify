package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceSavoringFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M24 1.5a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37a18.25 18.25 0 1 1 18.25-18.25A18.25 18.25 0 0 1 24 38.5Z"/><path fill="#fff48c" d="M18 5.5a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="#45413c" d="M8 45.5a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ffaa54" d="M38.5 26.5c0 .83-1.12 1.5-2.5 1.5s-2.5-.67-2.5-1.5S34.62 25 36 25s2.5.67 2.5 1.5Zm-29 0c0 .83 1.12 1.5 2.5 1.5s2.5-.67 2.5-1.5S13.38 25 12 25s-2.5.67-2.5 1.5Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M10 22.75a1.75 1.75 0 0 1 3.5 0m21 0a1.75 1.75 0 0 1 3.5 0M16.11 29.5a11.9 11.9 0 0 0 15.8 0"/><path fill="#ff6242" d="M31 30.24a12 12 0 0 1-6.91 2.26c3 2.49 2.22 6.25 5.23 5.89c4.12-.5 3.41-5.85 1.68-8.15Z"/><path fill="#ff866e" d="M29.84 38.29c3.53-.87 2.81-5.85 1.15-8a11.26 11.26 0 0 1-1.45.89c1.08 1.93 1.58 5.3.3 7.11Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M31 30.24a12 12 0 0 1-6.91 2.26c3 2.49 2.22 6.25 5.23 5.89c4.12-.5 3.41-5.85 1.68-8.15Z"/>`),
		g.Group(children),
	)
}