package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TapeDriveSolidAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<g id="clarityTapeDriveSolidAlerted0" fill="currentColor"><path d="M7.2 18a5 5 0 1 0 5-5a5 5 0 0 0-5 5Zm7 0a2 2 0 1 1-2-2a2 2 0 0 1 2.02 2Zm4.58 0a5 5 0 1 0 9.27-2.6h-5.82a3.71 3.71 0 0 1-2.17-.71A5 5 0 0 0 18.78 18Zm5-2a2 2 0 1 1-2 2a2 2 0 0 1 2.01-2Z"/><path d="M33.68 15.4H30V24H6V12h12.57a3.65 3.65 0 0 1 .48-2.11L21.29 6H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V15.38Z"/><path d="m26.85 1.14l-5.72 9.91a1.27 1.27 0 0 0 1.1 1.95h11.45a1.27 1.27 0 0 0 1.1-1.91l-5.72-9.95a1.28 1.28 0 0 0-2.21 0Z"/></g>`),
		g.Group(children),
	)
}