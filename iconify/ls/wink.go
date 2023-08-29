package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M358 0c198 0 359 155 359 349S556 698 358 698S0 543 0 349S160 0 358 0zm229 332l8-21c3-8-1-17-9-20l-117-42l108-62c7-4 9-13 5-20l-12-20c-4-7-13-9-19-5l-161 92c-7 4-8 10-7 23s3 18 10 21l175 63c7 3 16-2 19-9zM258 171h-41c-11 0-19 9-19 20v100c0 11 8 21 19 21h41c10 0 20-10 20-21V191c0-11-10-20-20-20zm-30 242l-39 13c-4 1-7 4-8 7c-2 3-2 7-1 11s36 105 178 104c142 1 178-100 179-104s1-8-1-11c-1-3-5-6-9-7l-34-12l-5-1c-6 0-11 5-13 10c0 1-22 61-117 61c-94 0-115-58-116-61c-2-5-8-10-14-10z"/>`),
		g.Group(children),
	)
}