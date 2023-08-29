package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RestorePageOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17.75q1.95 0 3.35-1.4t1.4-3.35q0-1.95-1.4-3.35T12 8.25q-.95 0-1.775.35t-1.475.95V8h-1.5v4.25h4.25v-1.5H9.7q.425-.45 1.025-.725T12 9.75q1.35 0 2.3.95t.95 2.3q0 1.35-.95 2.3t-2.3.95q-1.1 0-1.925-.638T8.9 14H7.35q.35 1.625 1.638 2.688T12 17.75ZM4 22V2h10l6 6v14H4Zm2-2h12V8.85L13.15 4H6v16Zm0 0V4v16Z"/>`),
		g.Group(children),
	)
}