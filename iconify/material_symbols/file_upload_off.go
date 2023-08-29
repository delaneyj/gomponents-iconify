package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileUploadOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 4l5 5l-1.4 1.45l-2.6-2.6v2.3l-2-2l-1.575-1.575L12 4Zm-1 12v-5l2 2v3h-2Zm8.775 6.625L17.15 20H6q-.825 0-1.413-.588T4 18v-3h2v3h9.15L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425ZM20 17.15l-2-2V15h2v2.15Z"/>`),
		g.Group(children),
	)
}