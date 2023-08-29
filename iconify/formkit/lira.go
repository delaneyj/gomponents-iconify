package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lira(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 15.5C3.86 15.5.5 12.14.5 8S3.86.5 8 .5s7.5 3.36 7.5 7.5s-3.36 7.5-7.5 7.5Zm0-14C4.42 1.5 1.5 4.42 1.5 8s2.92 6.5 6.5 6.5s6.5-2.92 6.5-6.5S11.58 1.5 8 1.5Z"/><path fill="currentColor" d="M7.5 13h-1c-.28 0-.5-.22-.5-.5v-9c0-.28.22-.5.5-.5s.5.22.5.5V12h.5A2.5 2.5 0 0 0 10 9.5c0-.28.22-.5.5-.5s.5.22.5.5c0 1.93-1.57 3.5-3.5 3.5Z"/><path fill="currentColor" d="M5 8.5c-.21 0-.4-.13-.47-.34c-.09-.26.05-.55.32-.63l4.5-1.5c.26-.09.55.05.63.32c.09.26-.05.55-.32.63l-4.5 1.5c-.05.02-.11.03-.16.03Zm0-2c-.21 0-.4-.13-.47-.34c-.09-.26.05-.55.32-.63l4.5-1.5c.26-.09.55.05.63.32c.09.26-.05.55-.32.63l-4.5 1.5c-.05.02-.11.03-.16.03Z"/>`),
		g.Group(children),
	)
}