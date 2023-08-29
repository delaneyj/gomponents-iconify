package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buromobelexperte(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v6h6V5H5zm8 0v6h6V5h-6zm8 0v6h6V5h-6zM7 7h2v2H7V7zm8 0h2v2h-2V7zm8 0h2v2h-2V7zM5 13v6h6v-6H5zm8 0v6h6v-6h-6zm8 0v6h6v-6h-6zM7 15h2v2H7v-2zm8 0h2v2h-2v-2zm8 0h2v2h-2v-2zM5 21v6h6v-6H5zm8 0v6h6v-6h-6zm8 0v6h6v-6h-6zM7 23h2v2H7v-2zm8 0h2v2h-2v-2z"/>`),
		g.Group(children),
	)
}