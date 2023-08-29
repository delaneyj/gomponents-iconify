package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="fePhone0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePhone1" fill="currentColor"><path id="fePhone2" d="M4.024 9L4 8.931C3.46 7.384 3 5.27 3 4c0-.55.45-1 1-1h3a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-.837A16.054 16.054 0 0 0 15 17.837V17a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v3c0 .45-.55 1-1 1c-1.725 0-3.44-.456-5-1c-5.114-1.832-9.168-5.886-10.976-11Z"/></g></g>`),
		g.Group(children),
	)
}