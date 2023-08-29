package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.303 21.707V8.275l4.48-4.42l-2.02-2.05l-4.127 4.07H8.76V2.084H5.883v3.793H1.8v2.877h4.083v15.832h15.542v4.61h2.878v-4.61H29.2v-2.878h-4.897zM19.72 8.753L8.76 19.565V8.753h10.96zm-9.032 12.953l10.735-10.59v10.59H10.69z"/>`),
		g.Group(children),
	)
}