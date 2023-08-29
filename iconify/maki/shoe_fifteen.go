package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoeFifteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9.5 7c-.414-.235-.89-.596-1.315-.948L6.01 3.222a.56.56 0 0 0-1 .278H5V5H3.209a.504.504 0 0 1-.357-.148S2.5 4 2 4h-.5a.5.5 0 0 0-.5.5V9h5.5c1.5 0 2 1 3.5 1h4v-.5C14 8 10.547 7.594 9.5 7z" fill="currentColor"/><path d="M9.5 11c-.632 0-1.047-.207-1.526-.447C7.456 10.293 6.868 10 6 10H1v1.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V11c.632 0 1.046.207 1.526.447c.519.26 1.106.553 1.974.553h4a.5.5 0 0 0 .5-.5V11H9.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}