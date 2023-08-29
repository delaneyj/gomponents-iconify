package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HashCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M6 10.25a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1Zm0 5.5a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1Z"/><path d="M10.833 6.004a1 1 0 0 1 .914 1.08l-1 12a1 1 0 1 1-1.994-.167l1-12a1 1 0 0 1 1.08-.913Zm5.5 0a1 1 0 0 1 .914 1.08l-1 12a1 1 0 0 1-1.993-.167l1-12a1 1 0 0 1 1.08-.913Z"/><path d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z"/></g>`),
		g.Group(children),
	)
}