package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaWifiFill0"><g id="evaWifiFill1"><g id="evaWifiFill2" fill="currentColor"><circle cx="12" cy="19" r="1"/><path d="M12 14a5 5 0 0 0-3.47 1.4a1 1 0 1 0 1.39 1.44a3.08 3.08 0 0 1 4.16 0a1 1 0 1 0 1.39-1.44A5 5 0 0 0 12 14Zm0-5a9 9 0 0 0-6.47 2.75A1 1 0 0 0 7 13.14a7 7 0 0 1 10.08 0a1 1 0 0 0 .71.3a1 1 0 0 0 .72-1.69A9 9 0 0 0 12 9Z"/><path d="M21.72 7.93a14 14 0 0 0-19.44 0a1 1 0 0 0 1.38 1.44a12 12 0 0 1 16.68 0a1 1 0 0 0 .69.28a1 1 0 0 0 .72-.31a1 1 0 0 0-.03-1.41Z"/></g></g></g>`),
		g.Group(children),
	)
}