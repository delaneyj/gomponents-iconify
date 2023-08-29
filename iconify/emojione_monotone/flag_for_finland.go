package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForFinland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm27.142 23.116H31.017V4.025C31.344 4.014 31.67 4 32 4c13.064 0 24.065 8.995 27.142 21.116zM17.25 8.215v16.901H4.858c1.803-7.102 6.327-13.127 12.392-16.901zM4.858 38.883H17.25v16.902c-6.065-3.774-10.59-9.8-12.392-16.902zM32 60c-.33 0-.656-.014-.983-.025V38.883h28.125C56.066 51.005 45.064 60 32 60z"/>`),
		g.Group(children),
	)
}