package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#FFF" d="M117.64 102.31H10.36c-2.74 0-3.83-1.71-3.83-3.83V29.52c0-2.12 1.71-3.83 3.83-3.83h107.27c2.12 0 3.83 1.71 3.83 3.83v68.96a3.815 3.815 0 0 1-3.82 3.83z"/><path fill="#6FBFF0" stroke="#6FBFF0" stroke-miterlimit="10" stroke-width=".5" d="m89.01 56.73l30.4-27.29c.61-.51 1.8-.78 1.8-.78s-1.39-1.52-3.22-.09L70.26 66.44A10.57 10.57 0 0 1 64 68.72a10.57 10.57 0 0 1-6.26-2.28L10.01 28.57c-1.83-1.43-3.22.09-3.22.09s1.19.28 1.8.78l30.04 26.97L7.81 99.57c-.53.59.92 1.55 1.59.9l33.66-40.09l12.37 11.1c2.44 2.19 5.5 3.32 8.58 3.39c3.08-.07 6.13-1.2 8.58-3.39l12.29-11.04l33 40.02c.67.65 2.12-.31 1.59-.9L89.01 56.73z"/><path fill="none" stroke="#6FBFF0" stroke-miterlimit="10" stroke-width="4" d="M117.64 102.31H10.36c-2.74 0-3.83-1.71-3.83-3.83V29.52c0-2.12 1.71-3.83 3.83-3.83h107.27c2.12 0 3.83 1.71 3.83 3.83v68.96a3.815 3.815 0 0 1-3.82 3.83z"/>`),
		g.Group(children),
	)
}