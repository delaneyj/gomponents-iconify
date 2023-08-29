package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm16.362-3.119L15.54 9.906L17.4 11.3l-1.2 1.6l-3.2-2.4V9.04l3.381-1.902l.98 1.743ZM9 8v4H7V8h2Zm-.1 5.634l.5.865C9.922 15.4 10.89 16 12 16s2.08-.601 2.6-1.5l.5-.866l1.731 1.001l-.5.866A4.998 4.998 0 0 1 12 18a4.998 4.998 0 0 1-4.33-2.5l-.501-.865L8.9 13.634Z"/>`),
		g.Group(children),
	)
}