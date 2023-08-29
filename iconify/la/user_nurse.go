package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserNurse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10.156 4L10 4.813l-1 5.625V14h.094a7.032 7.032 0 0 0 3 4.813C8.527 20.343 6 23.883 6 28h2a7.988 7.988 0 0 1 3-6.25v.656l.281.313l4 4l.719.687l.719-.687l4-4l.281-.313v-.656A7.988 7.988 0 0 1 24 28h2c0-4.117-2.527-7.656-6.094-9.188a7.032 7.032 0 0 0 3-4.812H23v-3.563l-1-5.624L21.844 4zm1.688 2h8.312L21 10.781V12H11v-1.219zM15 7v1h-1v2h1v1h2v-1h1V8h-1V7zm-3.906 7h9.812A5.001 5.001 0 0 1 16 18a5.001 5.001 0 0 1-4.906-4zM16 20a8.05 8.05 0 0 1 3 .563v1l-3 3l-3-3v-1A8.05 8.05 0 0 1 16 20z"/>`),
		g.Group(children),
	)
}