package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 36a92 92 0 1 0 92 92a92.1 92.1 0 0 0-92-92Zm0 176a84 84 0 1 1 84-84a84.09 84.09 0 0 1-84 84ZM58.83 26.83l-32 32a4 4 0 0 1-5.66-5.66l32-32a4 4 0 0 1 5.66 5.66Zm176 32a4 4 0 0 1-5.66 0l-32-32a4 4 0 0 1 5.66-5.66l32 32a4 4 0 0 1 0 5.66ZM188 128a4 4 0 0 1-4 4h-56a4 4 0 0 1-4-4V72a4 4 0 0 1 8 0v52h52a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}