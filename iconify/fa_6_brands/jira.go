package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jira(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M490 241.7C417.1 169 320.6 71.8 248.5 0C83 164.9 6 241.7 6 241.7c-7.9 7.9-7.9 20.7 0 28.7C138.8 402.7 67.8 331.9 248.5 512c379.4-378 15.7-16.7 241.5-241.7c8-7.9 8-20.7 0-28.6zm-241.5 90l-76-75.7l76-75.7l76 75.7l-76 75.7z"/>`),
		g.Group(children),
	)
}