package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleWithLeftHalfBlack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M37 44.286V7c-7.956 0-15.587 3.055-21.213 8.494C10.16 20.933 7 28.309 7 36c0 7.691 3.16 15.068 8.787 20.506C21.413 61.945 29.044 65 37 65V44.286Z"/><path fill="#fff" d="M36 45V9a28 28 0 1 1 0 56V45Z"/><path fill="#3F3F3F" d="M37 44.286V7c-7.956 0-15.587 3.055-21.213 8.494C10.16 20.933 7 28.309 7 36c0 7.691 3.16 15.068 8.787 20.506C21.413 61.945 29.044 65 37 65V44.286Z"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-linejoin="round" d="M36 64c15.464 0 28-12.536 28-28S51.464 8 36 8S8 20.536 8 36s12.536 28 28 28Z"/><path stroke-linecap="round" d="M36 8v53"/></g>`),
		g.Group(children),
	)
}