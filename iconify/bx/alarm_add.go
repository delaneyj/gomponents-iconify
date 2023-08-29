package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4c-4.879 0-9 4.121-9 9s4.121 9 9 9s9-4.121 9-9s-4.121-9-9-9zm0 16c-3.794 0-7-3.206-7-7s3.206-7 7-7s7 3.206 7 7s-3.206 7-7 7z"/><path fill="currentColor" d="M13 8h-2v4H7v2h4v4h2v-4h4v-2h-4zm7.292-1.292l-3.01-3l1.412-1.417l3.01 3zM5.282 2.294L6.7 3.706l-2.99 3l-1.417-1.413z"/>`),
		g.Group(children),
	)
}