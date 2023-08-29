package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StumbleuponWithCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 .4C4.698.4.4 4.698.4 10s4.298 9.6 9.6 9.6s9.6-4.298 9.6-9.6S15.302.4 10 .4zm0 7.385a.53.53 0 0 0-.531.529v3.168a2.262 2.262 0 0 1-4.522 0v-1.326h1.729v1.326a.53.53 0 0 0 .531.529a.53.53 0 0 0 .531-.529V8.314a2.262 2.262 0 0 1 4.523.001v.603l-1.04.334l-.69-.334v-.604A.53.53 0 0 0 10 7.785zm5.053 3.697a2.262 2.262 0 0 1-4.523 0v-1.354l.69.334l1.04-.334v1.354a.53.53 0 0 0 1.061 0v-1.326h1.731v1.326z"/>`),
		g.Group(children),
	)
}