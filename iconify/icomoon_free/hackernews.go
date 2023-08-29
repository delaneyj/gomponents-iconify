package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hackernews(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0v16h16V0H0zm8.5 9.125V12.5h-1V9.125L4.766 4H5.9L8 7.938L10.1 4h1.134L8.5 9.125z"/>`),
		g.Group(children),
	)
}