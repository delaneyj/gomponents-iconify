package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><g opacity=".2"><path d="M19.568 14.058a1 1 0 0 1-.054 1.721l-8.033 4.408A1 1 0 0 1 10 19.311V9.817a1 1 0 0 1 1.535-.845l8.033 5.086Z"/><path fill-rule="evenodd" d="M17.067 14.841L12 11.633v5.988l5.067-2.78Zm2.447.938a1 1 0 0 0 .054-1.721l-8.033-5.086A1 1 0 0 0 10 9.817v9.494a1 1 0 0 0 1.481.876l8.033-4.408Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="m9 18.321l9.014-4.883L9 7.804v10.517Zm9.49-4.003a1 1 0 0 0 .054-1.728L9.53 6.956A1 1 0 0 0 8 7.804v10.517a1 1 0 0 0 1.476.88l9.015-4.883Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}