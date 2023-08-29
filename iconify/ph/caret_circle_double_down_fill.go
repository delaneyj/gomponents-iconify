package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretCircleDoubleDownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M201.57 54.43A104.15 104.15 0 1 0 232 128a104.17 104.17 0 0 0-30.43-73.57Zm-35.9 95.24l-32 32a8 8 0 0 1-11.32 0l-32-32a8 8 0 1 1 11.32-11.32L128 164.71l26.35-26.36a8 8 0 1 1 11.32 11.32Zm0-56l-32 32a8 8 0 0 1-11.32 0l-32-32a8 8 0 1 1 11.32-11.32L128 108.68l26.35-26.36a8 8 0 1 1 11.32 11.32Z"/>`),
		g.Group(children),
	)
}