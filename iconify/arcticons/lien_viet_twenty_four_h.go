package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LienVietTwentyFourH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17.395 20.831A19.5 19.5 0 0 1 24 27.098m0 .001a19.499 19.499 0 0 1 6.604-6.267m.001-.001A19.5 19.5 0 0 1 24 14.564"/><path d="M23.999 14.565a19.499 19.499 0 0 1-6.604 6.267"/></g><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M17.395 20.831L4.5 20.854m26.105-.022l12.895.023"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M43.5 20.855a19.5 19.5 0 0 1-39 0"/><path d="M37.21 20.855a13.21 13.21 0 1 1-26.42 0m0 0a13.21 13.21 0 1 1 26.42 0"/></g><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M17.395 20.831h13.21"/>`),
		g.Group(children),
	)
}