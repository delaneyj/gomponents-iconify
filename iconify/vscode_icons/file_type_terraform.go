package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeTerraform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#813cf3" d="m12.042 6.858l8.029 4.59v9.014l-8.029-4.594v-9.01zM20.5 20.415l7.959-4.575V6.887L20.5 11.429v8.986zM3.541 11.01l8.03 4.589V6.59L3.541 2v9.01zm8.501 14.4L20.071 30v-9.043l-8.029-4.589v9.042z"/>`),
		g.Group(children),
	)
}