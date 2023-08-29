package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grids(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 21V3a1 1 0 0 0-1-1h-5v20h5a1 1 0 0 0 1-1Z" opacity=".5"/><path fill="currentColor" d="M1 3v18a1 1 0 0 0 1 1h5V2H2a1 1 0 0 0-1 1Z"/><path fill="currentColor" d="M9 2h6v20H9z" opacity=".5"/><path fill="currentColor" d="M7 2h2v20H7zm8 0h2v20h-2z" opacity=".25"/>`),
		g.Group(children),
	)
}