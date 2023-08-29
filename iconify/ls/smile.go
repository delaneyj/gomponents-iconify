package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M358 0c198 0 359 155 359 349S556 698 358 698S0 543 0 349S160 0 358 0zM258 171h-41c-11 0-19 9-19 20v100c0 11 8 21 19 21h41c10 0 20-10 20-21V191c0-11-10-20-20-20zm242 0h-41c-10 0-20 9-20 20v100c0 11 10 21 20 21h41c11 0 19-10 19-21V191c0-11-8-20-19-20zM228 413l-39 13c-4 1-7 4-8 7c-2 3-2 7-1 11s36 105 178 104c142 1 178-100 179-104s1-8-1-11c-1-3-5-6-9-7l-34-12l-5-1c-6 0-11 5-13 10c0 1-22 61-117 61c-94 0-115-58-116-61c-2-5-8-10-14-10z"/>`),
		g.Group(children),
	)
}