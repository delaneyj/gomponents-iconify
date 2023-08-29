package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceBlowingAKiss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M1.5 21.25a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M21.5 1.25a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37A18.25 18.25 0 1 1 39.75 20A18.25 18.25 0 0 1 21.5 38.25Z"/><path fill="#fff48c" d="M15.5 5.25a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/><path fill="#ebcb00" d="M39.39 30.19a20 20 0 0 1-8.19 8.53l-1.12-2.62l-.55-1.29a7.39 7.39 0 0 1 .16-5.81a5.71 5.71 0 0 1 3.9-3.32a4.46 4.46 0 0 1 4.62 1.69a4.89 4.89 0 0 1 .79 1.45Z"/><path fill="#45413c" d="M5.5 45.25a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M1.5 21.25a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ff6242" d="M46 29.63a4.64 4.64 0 0 1 .09 4.13A6 6 0 0 1 42.61 37l-8.25 2.77l-3.41-8a6 6 0 0 1 .12-4.77a4.65 4.65 0 0 1 3.15-2.68a3.65 3.65 0 0 1 4.34 2.57L39.2 29l2-.89a3.65 3.65 0 0 1 4.8 1.52Z"/><path fill="#ff866e" d="M32.51 28.16a3.64 3.64 0 0 1 4.34 2.56l.64 2.06l2-.89a3.64 3.64 0 0 1 4.8 1.55a4.3 4.3 0 0 1 .52 2.15a6.27 6.27 0 0 0 1.28-1.83a4.64 4.64 0 0 0-.09-4.13a3.65 3.65 0 0 0-4.8-1.55l-2 .89l-.64-2.06a3.65 3.65 0 0 0-4.34-2.57A4.65 4.65 0 0 0 31.07 27a6 6 0 0 0-.52 2.19a4.22 4.22 0 0 1 1.96-1.03Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M46 29.63a4.64 4.64 0 0 1 .09 4.13A6 6 0 0 1 42.61 37l-8.25 2.77l-3.41-8a6 6 0 0 1 .12-4.77a4.65 4.65 0 0 1 3.15-2.68a3.65 3.65 0 0 1 4.34 2.57L39.2 29l2-.89a3.65 3.65 0 0 1 4.8 1.52Z"/><path fill="#ffaa54" d="M7 26.25c0 .83 1.12 1.5 2.5 1.5s2.5-.67 2.5-1.5s-1.12-1.5-2.5-1.5s-2.5.67-2.5 1.5Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M30 21.5a1.75 1.75 0 0 1 3.5 0"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M9.5 20.25a1 1 0 1 0 1-1a1 1 0 0 0-1 1Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21.51 28a2 2 0 0 1 1-.27a2 2 0 1 1 0 4a2 2 0 1 1 0 4a2 2 0 0 1-1-.27"/>`),
		g.Group(children),
	)
}