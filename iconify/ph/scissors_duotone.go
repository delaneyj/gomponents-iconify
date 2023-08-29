package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M40.2 95.8a28 28 0 1 1 39.6 0a28 28 0 0 1-39.6 0Zm0 64.4a28 28 0 1 0 39.6 0a28 28 0 0 0-39.6 0Z" opacity=".2"/><path d="M157.73 113.13a8 8 0 0 1 2.09-11.13l67.66-46.3a8 8 0 0 1 9 13.21l-67.67 46.3a7.92 7.92 0 0 1-4.51 1.4a8 8 0 0 1-6.57-3.48Zm80.87 85.09a8 8 0 0 1-11.12 2.08L136 137.7l-42.51 29.08a36 36 0 1 1-9-13.19L121.83 128l-37.39-25.59a35.86 35.86 0 1 1 9-13.19l143 97.87a8 8 0 0 1 2.16 11.13ZM80 180a20 20 0 1 0-5.86 14.14A19.85 19.85 0 0 0 80 180Zm-5.86-89.87a20 20 0 1 0-28.28 0a19.85 19.85 0 0 0 28.28 0Z"/></g>`),
		g.Group(children),
	)
}