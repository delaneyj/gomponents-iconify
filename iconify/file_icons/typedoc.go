package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Typedoc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m234 3l232 125v254L233 509V257L0 129L234 3zM0 380.55l214.076-112.321L0 151.672V380.55z"/>`),
		g.Group(children),
	)
}