package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoftIceCreamOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M14.5 45.5a9.5 1.5 0 1 0 19 0a9.5 1.5 0 1 0-19 0Z" opacity=".15"/><path fill="#f0d5a8" d="M31 31.89H17a.9.9 0 0 1-.88-.69l-.91-3.63a.91.91 0 0 1 .88-1.13h15.82a.91.91 0 0 1 .88 1.13l-.91 3.63a.9.9 0 0 1-.88.69Z"/><path fill="#f7e5c6" d="M15.49 28.71h17l.28-1.14a.91.91 0 0 0-.88-1.13h-15.8a.91.91 0 0 0-.88 1.13Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M31 31.89H17a.9.9 0 0 1-.88-.69l-.91-3.63a.91.91 0 0 1 .88-1.13h15.82a.91.91 0 0 1 .88 1.13l-.91 3.63a.9.9 0 0 1-.88.69Z"/><path fill="#f0d5a8" d="M26.15 45.5h-4.3a2.72 2.72 0 0 1-2.7-2.36l-1.5-11.25h12.7l-1.5 11.25a2.72 2.72 0 0 1-2.7 2.36Z"/><path fill="#debb7e" d="m30.05 34.16l.3-2.27h-12.7l.3 2.27h12.1z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M26.15 45.5h-4.3a2.72 2.72 0 0 1-2.7-2.36l-1.5-11.25h12.7l-1.5 11.25a2.72 2.72 0 0 1-2.7 2.36Z"/><path fill="#fffef2" d="M33 16.79a3.13 3.13 0 0 0 .5-1.69a3.17 3.17 0 0 0-3.18-3.17h-1.13a2.71 2.71 0 0 0-2-4.54h-1.37a2.7 2.7 0 0 1-2.35-4.05a5.89 5.89 0 0 0-5.37 5.84a8.35 8.35 0 0 0 .46 2.75h-.91a3.17 3.17 0 0 0-3.18 3.17a3.13 3.13 0 0 0 .5 1.69a5 5 0 0 0 1.77 9.65h14.52A5 5 0 0 0 33 16.79Z"/><path fill="#fff" d="M14.47 23.27a5 5 0 0 1 3.22-4.66a3.16 3.16 0 0 1 2.68-4.87h.91a8.57 8.57 0 0 1-.46-2.74a5.89 5.89 0 0 1 2.68-4.94a2.63 2.63 0 0 1 0-2.75a5.89 5.89 0 0 0-5.4 5.87a8.35 8.35 0 0 0 .46 2.75h-.91a3.17 3.17 0 0 0-3.18 3.17a3.13 3.13 0 0 0 .5 1.69a5 5 0 0 0 .52 9.49a5 5 0 0 1-1.02-3.01Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33 16.79a3.13 3.13 0 0 0 .5-1.69a3.17 3.17 0 0 0-3.18-3.17h-1.13a2.71 2.71 0 0 0-2-4.54h-1.37a2.7 2.7 0 0 1-2.35-4.05a5.89 5.89 0 0 0-5.37 5.84a8.35 8.35 0 0 0 .46 2.75h-.91a3.17 3.17 0 0 0-3.18 3.17a3.13 3.13 0 0 0 .5 1.69a5 5 0 0 0 1.77 9.65h14.52A5 5 0 0 0 33 16.79Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M31.71 20.09v1.36a1.36 1.36 0 1 1-2.72 0v-3.17m-5.44-3.63v2.27a1.37 1.37 0 0 1-2.73 0V16"/><path fill="#fffef2" d="M19.92 25.54v2.27a1.37 1.37 0 0 1-2.73 0v-2.27Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M19.92 26.44v1.37a1.37 1.37 0 0 1-2.73 0v-2.27"/>`),
		g.Group(children),
	)
}