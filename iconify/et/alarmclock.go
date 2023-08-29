package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alarmclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M15 2C7.28 2 1 8.28 1 16s6.28 14 14 14s14-6.28 14-14S22.72 2 15 2zm0 27C7.832 29 2 23.168 2 16S7.832 3 15 3s13 5.832 13 13s-5.832 13-13 13z"/><path d="M7.5 0C3.364 0 0 3.364 0 7.5v1.062a.5.5 0 0 0 1 0V7.5C1 3.916 3.916 1 7.5 1h.875a.5.5 0 0 0 0-1H7.5zm14.875 0H21.5a.5.5 0 0 0 0 1h.875c3.584 0 6.5 2.916 6.5 6.5v1.062a.5.5 0 0 0 1 0V7.5c0-4.136-3.364-7.5-7.5-7.5zM15 16.293V9.319a.5.5 0 0 0-1 0V16.5c0 .133.053.26.146.354l4.914 4.914a.498.498 0 0 0 .708 0a.5.5 0 0 0 0-.707L15 16.293zM1.9 31.8l3-4a.5.5 0 0 0-.801-.599l-3 4a.5.5 0 0 0 .801.599zm24-4.6a.5.5 0 0 0-.801.599l3 4a.5.5 0 0 0 .801-.6l-3-3.999z"/></g>`),
		g.Group(children),
	)
}