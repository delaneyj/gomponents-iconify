package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Milestone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.586 6.586A1.986 1.986 0 0 0 23.17 6H16V2h-2v4H6a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h8v14h2V16h7.171a1.986 1.986 0 0 0 1.415-.586L29 11ZM23.17 14H6V8h17.172l3 3Z"/>`),
		g.Group(children),
	)
}