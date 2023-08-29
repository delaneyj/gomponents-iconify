package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenguinTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#656769" d="M5.51 22.01a18.49 18.49 0 1 0 36.98 0a18.49 18.49 0 1 0-36.98 0Z"/><path fill="#87898c" d="M24 7.22a18.5 18.5 0 0 1 18.4 16.64c.06-.61.09-1.23.09-1.85a18.49 18.49 0 0 0-37 0c0 .62 0 1.24.09 1.85A18.5 18.5 0 0 1 24 7.22Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M5.51 22.01a18.49 18.49 0 1 0 36.98 0a18.49 18.49 0 1 0-36.98 0Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M39 32.78a18.49 18.49 0 0 1-30 0a22.06 22.06 0 0 1-.83-6c0-7.59 3.66-13.21 8.19-13.21c3.52 0 6.47 1.46 7.66 5.81c1.19-4.29 4.15-5.81 7.66-5.81c4.52 0 8.18 5.62 8.18 13.21a21.62 21.62 0 0 1-.86 6Z"/><path fill="#45413c" d="M11.85 45.42a12.15 1.58 0 1 0 24.3 0a12.15 1.58 0 1 0-24.3 0Z" opacity=".15"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M30.87 21.48a1.58 1.58 0 1 0 3.16 0a1.58 1.58 0 1 0-3.16 0Zm-16.9 0a1.58 1.58 0 1 0 3.16 0a1.58 1.58 0 1 0-3.16 0Z"/><path fill="#ffaa54" d="M38.41 27.1c0 .55-.86 1-1.92 1s-1.92-.45-1.92-1s.86-1 1.92-1s1.92.45 1.92 1Zm-24.98 0c0 .55-.86 1-1.92 1s-1.92-.45-1.92-1s.86-1 1.92-1s1.92.45 1.92 1Z"/><path fill="#ebcb00" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M28.23 30.36a4.23 4.23 0 1 1-8.46 0a6.48 6.48 0 0 1 8.46 0Z"/><path fill="#ffe500" d="M28.37 30.3a1.05 1.05 0 0 0 .37-.85c-.2-3-2.25-5.33-4.74-5.33s-4.54 2.35-4.74 5.33a1.05 1.05 0 0 0 .37.85A6.86 6.86 0 0 0 24 32.05a6.86 6.86 0 0 0 4.37-1.75Z"/><path fill="#fff48c" d="M19.64 30.31A4.83 4.83 0 0 1 24 26.77a4.83 4.83 0 0 1 4.36 3.54a1.05 1.05 0 0 0 .37-.85c-.2-3-2.25-5.33-4.74-5.33s-4.54 2.35-4.74 5.33a1.05 1.05 0 0 0 .37.85Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M28.37 30.3a1.05 1.05 0 0 0 .37-.85c-.2-3-2.25-5.33-4.74-5.33s-4.54 2.35-4.74 5.33a1.05 1.05 0 0 0 .37.85A6.86 6.86 0 0 0 24 32.05a6.86 6.86 0 0 0 4.37-1.75Z"/>`),
		g.Group(children),
	)
}