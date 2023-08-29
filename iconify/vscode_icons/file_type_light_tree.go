package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeLightTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M3.021 2.022h1.997v27.955H3.021zM28.98 27.98H5.018v1.997H28.98zm-6.99-8.985H5.019v1.997h16.973zM13.006 9.01H5.018v1.997h7.987z"/>`),
		g.Group(children),
	)
}