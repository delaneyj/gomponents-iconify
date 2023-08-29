package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomanFrowningTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#a86c4d" d="M24 8.82a10.88 10.88 0 0 1 10.88 10.87v14.19a1.24 1.24 0 0 1-.77 1.12a27.09 27.09 0 0 1-20.22 0a1.24 1.24 0 0 1-.77-1.14V19.69A10.88 10.88 0 0 1 24 8.82Z"/><path fill="#de926a" d="M24 8.82a10.88 10.88 0 0 0-10.88 10.87v3a10.88 10.88 0 0 1 21.76 0v-3A10.88 10.88 0 0 0 24 8.82Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 8.82a10.88 10.88 0 0 1 10.88 10.87v14.19a1.24 1.24 0 0 1-.77 1.12h0a27.09 27.09 0 0 1-20.22 0h0a1.24 1.24 0 0 1-.77-1.14V19.69A10.88 10.88 0 0 1 24 8.82Z"/><path fill="#45413c" d="M10.5 45.54a13.5 1.5 0 1 0 27 0a13.5 1.5 0 1 0-27 0Z" opacity=".15"/><path fill="#ff87af" d="M24 31.22A11.88 11.88 0 0 1 35.88 43.1V45H12.12v-1.9A11.88 11.88 0 0 1 24 31.22Z"/><path fill="#ff6196" d="M24 31.22A11.88 11.88 0 0 0 12.12 43.1V45h.07A11.88 11.88 0 0 1 24 34.43A11.88 11.88 0 0 1 35.81 45h.07v-1.9A11.88 11.88 0 0 0 24 31.22Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 31.22h0A11.88 11.88 0 0 1 35.88 43.1V45h0h-23.76h0v-1.9A11.88 11.88 0 0 1 24 31.22ZM17.06 45v-2.41M30.94 45v-2.41"/><path fill="#ffb59e" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21.46 27.17h5.09v7.94h-5.09Z"/><path fill="#ffcebf" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33.07 22a8.77 8.77 0 0 1-5.68-5l-.62-1.49a15.92 15.92 0 0 1-9.95 6.18l-1.89.31a1.81 1.81 0 0 0 0 3.62h.12a9 9 0 0 0 17.9 0h.12a1.81 1.81 0 1 0 0-3.62Z"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M18.57 23.41a.77.77 0 1 0 .77-.77a.76.76 0 0 0-.77.77Zm10.86 0a.77.77 0 1 1-.77-.77a.76.76 0 0 1 .77.77Z"/><path fill="#ffb59e" d="M17.06 26.69a1 .6 0 1 0 2 0a1 .6 0 1 0-2 0Zm11.88 0a1 .6 0 1 0 2 0a1 .6 0 1 0-2 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21.65 29.87a2.73 2.73 0 0 1 4.7 0"/>`),
		g.Group(children),
	)
}