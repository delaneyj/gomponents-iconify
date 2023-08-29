package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M6.5 10.5V6a5.5 5.5 0 1 1 11 0v4.5M12 15v4m8.5 4.5h-17v-.25l.075-.327a26.342 26.342 0 0 0 0-11.846L3.5 10.75v-.25h17v.25l-.075.327a26.34 26.34 0 0 0 0 11.846l.075.327v.25Z"/>`),
		g.Group(children),
	)
}