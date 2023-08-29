package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotebookAndPen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTNotebookAndPen0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 6v36h26V6H4Z"/><path d="M12 42V6"/><path fill="#555" d="M44 6h-8v32l4 4l4-4V6Z"/><path d="M36 12h8M30 6H4m26 36H4M36 6v16m8-16v16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTNotebookAndPen0)"/>`),
		g.Group(children),
	)
}