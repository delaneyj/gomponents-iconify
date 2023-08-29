package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path d="M16.568 11.058a1 1 0 0 1-.054 1.721l-8.033 4.408A1 1 0 0 1 7 16.311V6.817a1 1 0 0 1 1.535-.845l8.033 5.086Z"/><path fill-rule="evenodd" d="M14.067 11.841L9 8.633v5.988l5.067-2.78Zm2.447.938a1 1 0 0 0 .054-1.721L8.535 5.972A1 1 0 0 0 7 6.817v9.494a1 1 0 0 0 1.481.876l8.033-4.408Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="m6 15.321l9.014-4.883L6 4.804v10.517Zm9.49-4.003a1 1 0 0 0 .054-1.728L6.53 3.956A1 1 0 0 0 5 4.804v10.517a1 1 0 0 0 1.476.88l9.015-4.883Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}