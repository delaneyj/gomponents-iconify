package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4c-4.878 0-9 4.122-9 9s4.122 9 9 9c4.879 0 9-4.122 9-9s-4.121-9-9-9zm5 10h-6V8h2v4h4v2zm3.292-7.292l-3.01-3l1.412-1.417l3.01 3zM5.282 2.294L6.7 3.706l-2.99 3l-1.417-1.413z"/>`),
		g.Group(children),
	)
}