package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeLightToml(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M22.76 6.83v3.25h-5v15.09h-3.5V10.08h-5V6.83Z"/><path fill="#7e7f7f" d="M2 2h6.2v3.09H5.34v21.8H8.2V30H2Zm28 28h-6.2v-3.09h2.86V5.11H23.8V2H30Z"/>`),
		g.Group(children),
	)
}