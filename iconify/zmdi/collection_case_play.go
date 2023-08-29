package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionCasePlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 176v235h341q0 17-12.5 29.5T341 453H43q-18 0-30.5-12.5T0 411V176h43zm320-85h106v234q0 18-12.5 30.5T427 368H128q-18 0-30.5-12.5T85 325V91h107V48q0-18 12.5-30.5T235 5h85q18 0 30.5 12.5T363 48v43zM235 48v43h85V48h-85zm0 256l117-85l-117-64v149z"/>`),
		g.Group(children),
	)
}