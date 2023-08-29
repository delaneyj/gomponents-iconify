package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Previous(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M21.125 0H4.875A4.874 4.874 0 0 0 0 4.875v16.25A4.874 4.874 0 0 0 4.875 26h16.25A4.874 4.874 0 0 0 26 21.125V4.875A4.874 4.874 0 0 0 21.125 0zM16 17.949a.964.964 0 0 1-.479.817a.986.986 0 0 1-.952.039L8.34 13.857c-.337-.296-.537-.494-.537-.857s.237-.575.537-.857l6.229-4.949a.99.99 0 0 1 .952.04c.29.173.479.484.479.816v9.899z"/>`),
		g.Group(children),
	)
}