package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileyXEyesBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm0 192a84 84 0 1 1 84-84a84.09 84.09 0 0 1-84 84Zm60.49-92.49a12 12 0 0 1-17 17L164 129l-7.51 7.52a12 12 0 0 1-17-17L147 112l-7.52-7.51a12 12 0 0 1 17-17L164 95l7.51-7.52a12 12 0 0 1 17 17L181 112Zm-72 17a12 12 0 0 1-17 0L92 129l-7.51 7.52a12 12 0 0 1-17-17L75 112l-7.52-7.51a12 12 0 0 1 17-17L92 95l7.51-7.52a12 12 0 0 1 17 17L109 112l7.52 7.51a12 12 0 0 1-.03 16.98ZM144 180a16 16 0 1 1-16-16a16 16 0 0 1 16 16Z"/>`),
		g.Group(children),
	)
}