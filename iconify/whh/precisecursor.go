package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Precisecursor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 576H672q-13 0-22.5-9.5T640 544v-64q0-13 9.5-22.5T672 448h320q13 0 22.5 9.5t9.5 22.5v64q0 13-9.5 22.5T992 576zm-448 448h-64q-13 0-22.5-9.5T448 992V672q0-13 9.5-22.5T480 640h64q13 0 22.5 9.5T576 672v320q0 13-9.5 22.5T544 1024zm-32-448q-27 0-45.5-18.5T448 512t18.5-45.5T512 448t45.5 18.5T576 512t-18.5 45.5T512 576zm32-192h-64q-13 0-22.5-9.5T448 352V32q0-13 9.5-22.5T480 0h64q13 0 22.5 9.5T576 32v320q0 13-9.5 22.5T544 384zM352 576H32q-13 0-22.5-9.5T0 544v-64q0-13 9.5-22.5T32 448h320q13 0 22.5 9.5T384 480v64q0 13-9.5 22.5T352 576z"/>`),
		g.Group(children),
	)
}