package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Suitcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 3c-1.094 0-2 .906-2 2v1H9V5H7v1H5c-1.094 0-2 .906-2 2v16c0 1.094.906 2 2 2h22c1.094 0 2-.906 2-2V8c0-1.094-.906-2-2-2h-2V5h-2v1h-3V5c0-1.094-.906-2-2-2zm0 2h4v1h-4zM5 8h22v16h-2V9h-2v15H9V9H7v15H5z"/>`),
		g.Group(children),
	)
}