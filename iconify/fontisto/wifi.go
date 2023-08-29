package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.807 0h-.045C10.572 0 4.942 2.4.752 6.319l.013-.012L0 7.02l3.862 3.826l.72-.66c3.201-2.952 7.494-4.763 12.21-4.763s9.009 1.81 12.222 4.774l-.012-.011l.72.66l3.862-3.826l-.765-.713A23.292 23.292 0 0 0 16.845 0h-.041h.002z"/><path fill="currentColor" d="M27.405 12.03A15.673 15.673 0 0 0 16.83 7.95h-.674l-.007.015a15.716 15.716 0 0 0-9.958 4.076l.013-.012l-.787.713l3.893 3.855l.72-.63c1.791-1.606 4.171-2.587 6.78-2.587s4.989.982 6.79 2.596l-.01-.008l.72.63l3.893-3.854z"/><path fill="currentColor" d="m16.815 24l5.475-5.415l-.87-.713a7.022 7.022 0 0 0-4.625-1.5h.013h-.066a7.603 7.603 0 0 0-4.567 1.515l.02-.014l-.862.713l.795.787l3.96 3.915z"/>`),
		g.Group(children),
	)
}