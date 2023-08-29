package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmWatsonSpeechToText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 18H4c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h2.4l1.7 3l1.7-1l-2.3-4H4v-6h10v6h-3v2h3c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2zm5 8h2c3.9 0 7-3.1 7-7v-2h-2v2c0 2.8-2.2 5-5 5h-2v2zm-1-15h6v2h-6zm0-4h12v2H18zm0-4h12v2H18zM4 14h2v-2c0-2.8 2.2-5 5-5h4V5h-4c-3.9 0-7 3.1-7 7v2z"/>`),
		g.Group(children),
	)
}