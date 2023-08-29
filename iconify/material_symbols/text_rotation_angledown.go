package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotationAngledown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 21v-2h1.6L2.3 9.7l1.4-1.4l9.3 9.3V16h2v5h-5Zm5.1-5.55l-1.35-1.35l1.3-2.65L11.9 8.3L9.25 9.55L7.9 8.2l10.7-4.85l1.4 1.4l-4.9 10.7Zm-1.65-7.9l2.3 2.35l2.1-4.35l-.05-.05l-4.35 2.05Z"/>`),
		g.Group(children),
	)
}