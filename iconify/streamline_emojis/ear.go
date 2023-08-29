package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffda8f" d="M25.7 3.16a10.59 10.59 0 0 0-10.59 10.58v20.78a5 5 0 0 0 5 5a6.15 6.15 0 0 0 5.6-3.61a12.11 12.11 0 0 1 4-4.84a19.52 19.52 0 0 0 8.1-15.84A12.06 12.06 0 0 0 25.7 3.16Z"/><path fill="#ffbe3d" d="M25.86 11.5a7.72 7.72 0 0 1 7.58 6.29a18.69 18.69 0 0 0 .14-2.11a7.72 7.72 0 0 0-7.72-7.72a6.58 6.58 0 0 0-6.58 6.57v3.58a6.58 6.58 0 0 1 6.58-6.61Z"/><path fill="#45413c" d="M10.3 45.21a13.7 1.79 0 1 0 27.4 0a13.7 1.79 0 1 0-27.4 0Z" opacity=".15"/><path fill="#ffe9bd" d="M25.7 3.16a10.59 10.59 0 0 0-10.59 10.58v3.65A10.59 10.59 0 0 1 25.7 6.81a12 12 0 0 1 11.93 10.47a17.7 17.7 0 0 0 .12-2.07A12.06 12.06 0 0 0 25.7 3.16Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M25.7 3.16a10.59 10.59 0 0 0-10.59 10.58v20.78a5 5 0 0 0 5 5a6.15 6.15 0 0 0 5.6-3.61a12.11 12.11 0 0 1 4-4.84a19.52 19.52 0 0 0 8.1-15.84A12.06 12.06 0 0 0 25.7 3.16Z"/><path fill="#ffbe3d" d="M21.66 19.84a2.39 2.39 0 0 0-2.38 2.38v2.24a2.39 2.39 0 1 1 4.77 0v-2.24a2.39 2.39 0 0 0-2.39-2.38Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21.66 19.84a2.39 2.39 0 0 1 2.39 2.38v2.39a2.39 2.39 0 0 1-4.77 0v-2.39a2.39 2.39 0 0 1 2.38-2.38Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M19.28 22.22V14.5a6.58 6.58 0 0 1 6.58-6.57a7.72 7.72 0 0 1 7.72 7.72a18.16 18.16 0 0 1-1.79 7.76"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24.64 14.48c2.49 0 4.17 1 4.17 4.76"/>`),
		g.Group(children),
	)
}