package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h12q.825 0 1.413.588T20 4v16q0 .825-.588 1.413T18 22H6Zm0-2h12V4h-2v6.125q0 .3-.25.438t-.5-.013l-1.225-.75q-.25-.15-.513-.15T13 9.8l-1.225.75q-.25.15-.512.013T11 10.124V4H6v16Zm0 0V4v16Zm5-9.875q0 .3.263.438t.512-.013L13 9.8q.25-.15.512-.15t.513.15l1.225.75q.25.15.5.013t.25-.438q0 .3-.25.438t-.5-.013l-1.225-.75q-.25-.15-.513-.15T13 9.8l-1.225.75q-.25.15-.513.013T11 10.124Z"/>`),
		g.Group(children),
	)
}