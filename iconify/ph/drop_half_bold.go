package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropHalfBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M134.88 6.17a12 12 0 0 0-13.76 0a259 259 0 0 0-42.18 39C50.85 77.43 36 111.63 36 144a92 92 0 0 0 184 0c0-77.36-81.64-135.4-85.12-137.83ZM194.08 160H140v-16h56a68 68 0 0 1-1.92 16ZM140 120v-16h47a115 115 0 0 1 5.68 16Zm19.3-58.71A197.29 197.29 0 0 1 173.68 80H140V41.46a243.5 243.5 0 0 1 19.3 19.83ZM60 144c0-33.31 20-63.37 36.7-82.71A243.5 243.5 0 0 1 116 41.46v169.46A68.1 68.1 0 0 1 60 144Zm80 66.92V184h42.94A68 68 0 0 1 140 210.92Z"/>`),
		g.Group(children),
	)
}