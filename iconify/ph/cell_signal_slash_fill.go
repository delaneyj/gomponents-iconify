package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellSignalSlashFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M213.38 221.92a8 8 0 0 1-11.3-.54l-4.89-5.38H32a16 16 0 0 1-11.32-27.32l79.45-79.44l-58.06-63.86a8 8 0 1 1 11.84-10.77l160 176a8 8 0 0 1-.53 11.31ZM194.08 165a8 8 0 0 0 13.92-5.37V40a16 16 0 0 0-27.31-11.3l-52.51 52.5a8 8 0 0 0-.26 11Z"/>`),
		g.Group(children),
	)
}