package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSlashFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M53.92 34.62a8 8 0 1 0-11.84 10.76l39.24 43.17l-.06.12A64 64 0 1 0 72 216h88a87.34 87.34 0 0 0 31.8-5.93l10.28 11.31a8 8 0 1 0 11.84-10.76ZM160 40a87.35 87.35 0 0 0-52.11 17.08a8 8 0 0 0-1.18 11.83l107.05 117.75a8 8 0 0 0 5.92 2.62a8 8 0 0 0 5.94-2.67A88 88 0 0 0 160 40Z"/>`),
		g.Group(children),
	)
}