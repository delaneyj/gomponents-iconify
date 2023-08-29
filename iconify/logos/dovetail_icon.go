package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DovetailIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#230078" d="m256 182.857l-64 36.572v-73.143l-64-36.974l64-36.17l64 36.572v73.143Zm-128.01-73.143l-64 36.572V73.143L0 36.57L64 0l63.99 36.571v73.143Zm0 146.286l-64 36.571l.01-73.142l-64-36.572l63.99-36.571l64 36.571V256Z"/>`),
		g.Group(children),
	)
}