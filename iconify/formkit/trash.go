package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.5 4h-13c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h13c.28 0 .5.22.5.5s-.22.5-.5.5Z"/><path fill="currentColor" d="m11.02 3.81l-.44-1.46a.504.504 0 0 0-.48-.36H5.9a.5.5 0 0 0-.48.36l-.44 1.46l-.96-.29l.44-1.46C4.65 1.42 5.23.99 5.9.99h4.2c.67 0 1.24.43 1.44 1.07l.44 1.46l-.96.29Z"/><path fill="currentColor" d="M11.53 15H4.47c-.81 0-1.47-.64-1.5-1.45l-.34-9.87l1-.03l.34 9.87c0 .27.23.48.5.48h7.07c.27 0 .49-.21.5-.48l.34-9.87l1 .03l-.34 9.87c-.03.81-.69 1.45-1.5 1.45Z"/><path fill="currentColor" d="M6.5 11.62c-.28 0-.5-.22-.5-.5v-4c0-.28.22-.5.5-.5s.5.22.5.5v4c0 .28-.22.5-.5.5Zm3 0c-.28 0-.5-.22-.5-.5v-4c0-.28.22-.5.5-.5s.5.22.5.5v4c0 .28-.22.5-.5.5Z"/>`),
		g.Group(children),
	)
}