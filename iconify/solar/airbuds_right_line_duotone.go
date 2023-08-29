package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirbudsRightLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M21 18.667v.833c0 .14 0 .209-.003.268a2.35 2.35 0 0 1-2.23 2.229C18.71 22 18.64 22 18.5 22s-.209 0-.268-.003a2.35 2.35 0 0 1-2.229-2.23C16 19.71 16 19.64 16 19.5v-.833m5 0V7.559A5.588 5.588 0 0 0 15.44 2h-.104A3.353 3.353 0 0 0 12 5.336V8a3 3 0 0 0 3 3a1 1 0 0 1 1 1v6.667m5 0h-5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.676" d="M15 5v3" opacity=".5"/><circle cx="7.5" cy="16.5" r="5.5" stroke="currentColor" stroke-width="1.5" opacity=".5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M9 5.1A5.006 5.006 0 0 0 5.1 9" opacity=".5"/><path fill="currentColor" d="M6 14v-.75a.75.75 0 0 0-.75.75H6Zm-.75 5a.75.75 0 0 0 1.5 0h-1.5ZM6 16.5h-.75v.75H6v-.75Zm2.357 2.886a.75.75 0 1 0 1.286-.772l-1.286.772ZM6 14.75h1.5v-1.5H6v1.5Zm.75 1.75V14h-1.5v2.5h1.5Zm1.5-1.25c0 .154-.202.5-.75.5v1.5c1.109 0 2.25-.773 2.25-2h-1.5Zm-.75-.5c.548 0 .75.346.75.5h1.5c0-1.227-1.141-2-2.25-2v1.5ZM5.25 14v5h1.5v-5h-1.5Zm2.25 1.75H6v1.5h1.5v-1.5Zm-.643 1.136l1.5 2.5l1.286-.772l-1.5-2.5l-1.286.772Z"/></g>`),
		g.Group(children),
	)
}