package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactsBookUploadFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2v20H3V2h4Zm12.005 0C20.107 2 21 2.898 21 3.99v16.02c0 1.099-.893 1.99-1.995 1.99H9V2h10.005ZM15 8l-4 4h3v4h2v-4h3l-4-4Zm9 4v4h-2v-4h2Zm0-6v4h-2V6h2Z"/>`),
		g.Group(children),
	)
}