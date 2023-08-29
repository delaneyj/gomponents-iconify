package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiabetesMeasureOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m14.956 8.211l18-.015l.014 18l-18 .015l-.014-18Zm13.008 9.847c.002 2.133-1.757 3.81-3.997 3.812c-2.24.002-4.001-1.673-4.003-3.806c-.002-2.133 3.994-6.86 3.994-6.86s4.004 4.873 4.006 6.854Z" clip-rule="evenodd"/><path d="M16.984 27.986a1 1 0 0 0 .001 2l4-.003a1 1 0 1 0-.001-2l-4 .003Zm9.001.993a1 1 0 0 1 .999-1.001l4-.003a1 1 0 1 1 .001 2l-4 .003a1 1 0 0 1-1-1Zm-2.003 11.225a2 2 0 1 0-.003-4a2 2 0 0 0 .003 4Z"/><path fill-rule="evenodd" d="m12.978 35.213l5-.005l.007 9l12-.01l-.007-9l5-.003a3 3 0 0 0 2.997-3.003l-.02-25a3 3 0 0 0-3.002-2.997l-22 .018a3 3 0 0 0-2.998 3.002l.02 25a3 3 0 0 0 3.003 2.998ZM34.954 6.195l-22 .018a1 1 0 0 0-.999 1l.02 25a1 1 0 0 0 1.001 1l5-.005l-.001-2l12-.01l.001 2l5-.003a1 1 0 0 0 1-1.001l-.02-25a1 1 0 0 0-1.002-1ZM27.976 33.2l-8 .007l.008 9l8-.007l-.008-9Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}