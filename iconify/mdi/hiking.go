package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hiking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.47 8.67H19V23h-1.53V12.6c-.8-.16-1.55-.46-2.26-.89s-1.31-.93-1.82-1.51l-.62 3.07l2.23 2.2V23h-2v-6l-2.24-2.2L8.89 23H6.73S9.86 7.22 9.89 7.09c.11-.48.33-.85.7-1.09c.37-.27.74-.4 1.12-.4c.39 0 .75.09 1.08.27c.34.17.6.42.79.74l1.06 1.63c.29.54.68 1.01 1.17 1.39s1.05.67 1.66.87V8.67M8.55 5.89L7.4 5.65c-.57-.15-1.09-.03-1.56.29c-.46.32-.74.76-.84 1.34l-.81 3.98c-.03.29.03.55.19.79s.37.37.62.41l2.21.43l1.34-7M13 1c-1.1 0-2 .9-2 2s.9 2 2 2s2-.89 2-2s-.89-2-2-2Z"/>`),
		g.Group(children),
	)
}