package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadTrackTwoBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M15.75 21.273A9.971 9.971 0 0 1 12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10a9.985 9.985 0 0 1-.547 3.27l-1.863-1.86a2.25 2.25 0 0 0-3.182 0l-2.5 2.5a2.25 2.25 0 0 0 1.841 3.827v1.537Z" opacity=".5"/><path fill-rule="evenodd" d="M13.75 7a.75.75 0 0 0-1.5 0v5.55a2.75 2.75 0 1 0 1.5 2.45v-5c.627.471 1.406.75 2.25.75a.75.75 0 0 0 0-1.5A2.25 2.25 0 0 1 13.75 7Zm3.72 7.47a.75.75 0 0 1 1.06 0l2.5 2.5a.75.75 0 1 1-1.06 1.06l-1.22-1.22V22a.75.75 0 0 1-1.5 0v-5.19l-1.22 1.22a.75.75 0 1 1-1.06-1.06l2.5-2.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}