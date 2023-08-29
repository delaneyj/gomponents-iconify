package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmilingFaceWithHalo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M13 45.5a11 1.5 0 1 0 22 0a11 1.5 0 1 0-22 0Z" opacity=".15"/><path fill="#ffe500" d="M10.04 26.96a13.96 13.96 0 1 0 27.92 0a13.96 13.96 0 1 0-27.92 0Z"/><path fill="#ebcb00" d="M24 13a14 14 0 1 0 14 14a14 14 0 0 0-14-14Zm0 25.83a12.74 12.74 0 1 1 12.74-12.74A12.74 12.74 0 0 1 24 38.83Z"/><path fill="#fff48c" d="M19.81 15.79a4.19 1.05 0 1 0 8.38 0a4.19 1.05 0 1 0-8.38 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M10.04 26.96a13.96 13.96 0 1 0 27.92 0a13.96 13.96 0 1 0-27.92 0Z"/><path fill="#ffaa54" d="M34.12 30.45c0 .58-.78 1.05-1.74 1.05s-1.75-.47-1.75-1.05s.78-1 1.75-1s1.74.43 1.74 1Zm-20.24 0c0 .58.78 1.05 1.74 1.05s1.75-.47 1.75-1.05s-.78-1-1.75-1s-1.74.43-1.74 1Z"/><path fill="#fff48c" d="M24 1.5c-7.17 0-13 1.88-13 4.19s5.82 4.19 13 4.19S37 8 37 5.69S31.17 1.5 24 1.5Zm0 5.86c-5.09 0-9.22-.75-9.22-1.67S18.91 4 24 4s9.22.75 9.22 1.68S29.09 7.36 24 7.36Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M14.23 27.84a1.22 1.22 0 1 1 2.44 0m14.66 0a1.22 1.22 0 1 1 2.44 0"/><path fill="#ffb0ca" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M19.77 32.35a.66.66 0 0 0-.5.23a.67.67 0 0 0-.15.53a4.94 4.94 0 0 0 9.76 0a.67.67 0 0 0-.15-.53a.66.66 0 0 0-.5-.23Z"/><path fill="#ff87af" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M27.89 35.38a7.93 7.93 0 0 0-3.89-.91a7.93 7.93 0 0 0-3.89.91a4.92 4.92 0 0 0 7.78 0Z"/><path fill="#fffacf" d="M24 7.36c-5.09 0-9.22-.75-9.22-1.67c0 1.63 4.13 3 9.22 3s9.22-1.33 9.22-3c0 .92-4.13 1.67-9.22 1.67Z"/><path fill="#ffe500" d="M24 2.73c-5.09 0-9.22 1.33-9.22 3C14.78 4.76 18.91 4 24 4s9.22.75 9.22 1.68c0-1.62-4.13-2.95-9.22-2.95Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 1.5c-7.17 0-13 1.88-13 4.19s5.82 4.19 13 4.19S37 8 37 5.69S31.17 1.5 24 1.5Zm0 5.86c-5.09 0-9.22-.75-9.22-1.67S18.91 4 24 4s9.22.75 9.22 1.68S29.09 7.36 24 7.36Z"/>`),
		g.Group(children),
	)
}