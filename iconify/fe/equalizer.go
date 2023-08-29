package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equalizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feEqualizer0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feEqualizer1" fill="currentColor"><path id="feEqualizer2" d="M13.17 7a3.001 3.001 0 0 1 5.66 0H21a1 1 0 0 1 0 2h-2.17a3.001 3.001 0 0 1-5.66 0H3a1 1 0 1 1 0-2h10.17Zm-8 8a3.001 3.001 0 0 1 5.66 0H21a1 1 0 0 1 0 2H10.83a3.001 3.001 0 0 1-5.66 0H3a1 1 0 0 1 0-2h2.17ZM16 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-8 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`),
		g.Group(children),
	)
}