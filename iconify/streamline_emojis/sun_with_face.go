package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunWithFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M12.5 46.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/><path fill="#ffe500" d="M10.5 24a13.5 13.5 0 1 0 27 0a13.5 13.5 0 1 0-27 0Z"/><path fill="#ebcb00" d="M25.28 10.56a12.5 12.5 0 1 1-2.56 0a13.5 13.5 0 1 0 2.56 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M10.5 24a13.5 13.5 0 1 0 27 0a13.5 13.5 0 1 0-27 0Z"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M25.5 3a1.5 1.5 0 0 0-3 0v3a1.5 1.5 0 0 0 3 0ZM10.21 8.09a1.5 1.5 0 0 0-2.12 2.12l2.12 2.12a1.5 1.5 0 0 0 2.12-2.12ZM3 22.5a1.5 1.5 0 0 0 0 3h3a1.5 1.5 0 0 0 0-3Zm5.09 15.29a1.5 1.5 0 0 0 2.12 2.12l2.12-2.12a1.5 1.5 0 0 0-2.12-2.12ZM22.5 45a1.5 1.5 0 0 0 3 0v-3a1.5 1.5 0 0 0-3 0Zm15.29-5.09a1.5 1.5 0 0 0 2.12-2.12l-2.12-2.12a1.5 1.5 0 0 0-2.12 2.12ZM45 25.5a1.5 1.5 0 0 0 0-3h-3a1.5 1.5 0 0 0 0 3Zm-5.09-15.29a1.5 1.5 0 0 0-2.12-2.12l-2.12 2.12a1.5 1.5 0 0 0 2.12 2.12Z"/><path fill="#ffaa54" d="M14 27c0 .55.67 1 1.5 1s1.5-.45 1.5-1s-.67-1-1.5-1s-1.5.45-1.5 1Z"/><path fill="#ffb0ca" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M27.87 29.5a1 1 0 0 1 1 1.25a5 5 0 0 1-9.68 0a1 1 0 0 1 1-1.25Z"/><path fill="#ff87af" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 34.5a5 5 0 0 0 3.94-1.94A7.34 7.34 0 0 0 24 31.5a7.34 7.34 0 0 0-3.94 1.06A5 5 0 0 0 24 34.5Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M15 25a1.5 1.5 0 0 1 3 0"/><path fill="#ffaa54" d="M34 27c0 .55-.67 1-1.5 1s-1.5-.45-1.5-1s.67-1 1.5-1s1.5.45 1.5 1Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33 25a1.5 1.5 0 0 0-3 0"/><path fill="#fff48c" d="M20 13.5a4 1 0 1 0 8 0a4 1 0 1 0-8 0Z"/>`),
		g.Group(children),
	)
}