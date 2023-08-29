package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuroOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd" opacity=".2"><path d="M12.989 4.5C9.93 4.5 7.5 6.713 7.5 10c0 3.386 2.527 6 5.489 6c.743 0 1.451-.161 2.098-.454a1 1 0 1 1 .826 1.821a7.061 7.061 0 0 1-2.924.633C8.783 18 5.5 14.345 5.5 10c0-4.445 3.38-7.5 7.489-7.5c1.237 0 2.428.393 3.574 1.174a1 1 0 1 1-1.126 1.652c-.857-.584-1.663-.826-2.448-.826Z"/><path d="M3.5 8.5a1 1 0 0 1 1-1h9a1 1 0 1 1 0 2h-9a1 1 0 0 1-1-1Zm0 3.5a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2h-8a1 1 0 0 1-1-1Z"/></g><path fill-rule="evenodd" d="M11.989 3C8.668 3 6 5.423 6 9c0 3.626 2.716 6.5 5.989 6.5c.817 0 1.595-.177 2.305-.499a.5.5 0 0 1 .412.91a6.562 6.562 0 0 1-2.717.589C8.094 16.5 5 13.106 5 9c0-4.155 3.142-7 6.989-7c1.124 0 2.219.355 3.293 1.087a.5.5 0 1 1-.564.826c-.93-.633-1.83-.913-2.73-.913Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M3 7.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5ZM3 11a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 0 1h-8A.5.5 0 0 1 3 11Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}