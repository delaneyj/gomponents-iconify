package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApkExplorerAndEditorAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.41 21.365l-.544 2.793a8.544 8.544 0 0 0-2.49 1.429l-2.69-.925l-2.593 4.474l2.155 1.868c-.17.94-.17 1.895 0 2.86l-2.155 1.867l2.593 4.472l2.687-.923a8.475 8.475 0 0 0 2.493 1.428l.544 2.792h5.18l.533-2.765c.965-.404 1.755-.846 2.501-1.456l2.69.925l2.59-4.473l-2.152-1.87a7.94 7.94 0 0 0 0-2.855l2.154-1.872l-2.592-4.472l-2.687.923c-.667-.546-1.546-1.095-2.493-1.427l-.544-2.792h-5.18Zm2.543 7.67c1.905 0 3.45 1.54 3.45 3.441a3.447 3.447 0 0 1-3.45 3.443h0a3.447 3.447 0 0 1-3.452-3.442v0a3.447 3.447 0 0 1 3.452-3.442Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M19.556 15.55a1.53 1.53 0 1 1 0-3.06a1.53 1.53 0 0 1 0 3.06h0Zm8.905 0a1.53 1.53 0 1 1 1.53-1.53h0a1.53 1.53 0 0 1-1.53 1.53Z"/><path d="M24 6.13h0c6.024 0 10.907 4.883 10.907 10.907v1.556H13.094v-1.556C13.094 11.013 17.977 6.13 24 6.13ZM14.582 4.5l3.098 3.647M33.418 4.5L30.32 8.147"/></g>`),
		g.Group(children),
	)
}