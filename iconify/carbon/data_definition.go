package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataDefinition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="14" cy="14" r="2" fill="currentColor"/><path fill="currentColor" d="M20 30a.997.997 0 0 1-.707-.293L8.586 19A2.013 2.013 0 0 1 8 17.586V10a2.002 2.002 0 0 1 2-2h7.586A1.986 1.986 0 0 1 19 8.586l10.707 10.707a1 1 0 0 1 0 1.414l-9 9A.997.997 0 0 1 20 30ZM10 10v7.586l10 10L27.586 20l-10-10Z"/><path fill="currentColor" d="M12 30H4a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v8h-2V4H4v24h8Z"/>`),
		g.Group(children),
	)
}