package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeWithArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<g fill="#fcc21b" fill-rule="evenodd" clip-rule="evenodd"><path d="M14.85 69.89v51.42h98.3V69.89L64 94.33z"/><path d="M14.85 48.36v13.68L64 86.47l49.15-24.43V48.36z"/></g><path fill="#ed6c30" d="M63.85 62.28L40.24 30.06h14.73V0h17.62v30.06h14.88L63.85 62.28z"/>`),
		g.Group(children),
	)
}