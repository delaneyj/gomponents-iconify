package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdhesivePlasterLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M3.213 9.071a4.142 4.142 0 0 1 5.858-5.858L20.787 14.93a4.142 4.142 0 0 1-5.858 5.858L3.213 9.07Z"/><path stroke="currentColor" stroke-width="1.5" d="m12 17.858l-2.929 2.929a4.142 4.142 0 0 1-5.858-5.858L6.143 12L12 17.858Zm0-11.716l2.929-2.929a4.142 4.142 0 0 1 5.858 5.858L17.857 12L12 6.142Z" opacity=".5"/><path fill="currentColor" d="M15.841 12.871a.787.787 0 1 1-1.113 1.114a.787.787 0 0 1 1.113-1.114Zm-3.712-3.712a.788.788 0 1 0-1.114 1.114a.788.788 0 0 0 1.114-1.114Zm1.856 5.569a.788.788 0 1 1-1.114 1.113a.788.788 0 0 1 1.114-1.113Zm-3.712-3.713a.788.788 0 1 0-1.114 1.114a.788.788 0 0 0 1.114-1.114Zm6.497 4.641a.788.788 0 1 1-1.114 1.113a.788.788 0 0 1 1.114-1.113ZM9.345 8.23A.788.788 0 1 0 8.23 9.345a.788.788 0 0 0 1.114-1.113Z"/><path fill="currentColor" d="M13.057 11.944a.787.787 0 1 1-1.113 1.113a.787.787 0 0 1 1.113-1.113Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}