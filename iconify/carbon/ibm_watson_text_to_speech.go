package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmWatsonTextToSpeech(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 26h-2c-3.9 0-7-3.1-7-7v-2h2v2c0 2.8 2.2 5 5 5h2v2zm5-8h10c1.1 0 2 .9 2 2v6c0 1.1-.9 2-2 2h-2.4l-1.7 3l-1.7-1l2.3-4H28v-6H18v6h3v2h-3c-1.1 0-2-.9-2-2v-6c0-1.1.9-2 2-2zm10-4h-2v-2c0-2.8-2.2-5-5-5h-4V5h4c3.9 0 7 3.1 7 7v2zM2 11h6v2H2zm0-4h12v2H2zm0-4h12v2H2z"/>`),
		g.Group(children),
	)
}