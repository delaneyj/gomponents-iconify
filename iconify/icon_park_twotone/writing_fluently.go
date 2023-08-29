package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WritingFluently(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTWritingFluently0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" fill-rule="evenodd" d="M24 24v-5L39 4l5 5l-15 15h-5Z" clip-rule="evenodd"/><path d="M16 24H9a5 5 0 0 0 0 10h30a5 5 0 0 1 0 10H18"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTWritingFluently0)"/>`),
		g.Group(children),
	)
}