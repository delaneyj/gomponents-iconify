package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchAdvanced(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 6h-4V2h-2v4h-4v2h4v4h2V8h4V6zm-6 22.586l-5.975-5.975a9.023 9.023 0 1 0-1.414 1.414L22.586 30zM4 17a7 7 0 1 1 7 7a7.008 7.008 0 0 1-7-7z"/>`),
		g.Group(children),
	)
}