package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webdriverio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.875 0A1.87 1.87 0 0 0 0 1.875v20.25A1.87 1.87 0 0 0 1.875 24h20.25A1.87 1.87 0 0 0 24 22.125V1.875A1.87 1.87 0 0 0 22.125 0Zm.375 6H3v12h-.75Zm7.085 0h.79L5.29 18h-.791Zm6.79 0a6 6 0 1 1 0 12a6 6 0 0 1 0-12zm0 .75a5.25 5.25 0 1 0 .001 10.501a5.25 5.25 0 0 0-.001-10.501z"/>`),
		g.Group(children),
	)
}