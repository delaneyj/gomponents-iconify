package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaSpeakerFill0"><g id="evaSpeakerFill1"><g id="evaSpeakerFill2" fill="currentColor"><circle cx="12" cy="15.5" r="1.5"/><circle cx="12" cy="8" r="1"/><path d="M17 2H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3Zm-5 3a3 3 0 1 1-3 3a3 3 0 0 1 3-3Zm0 14a3.5 3.5 0 1 1 3.5-3.5A3.5 3.5 0 0 1 12 19Z"/></g></g></g>`),
		g.Group(children),
	)
}