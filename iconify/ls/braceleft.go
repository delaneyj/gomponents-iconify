package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Braceleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M289 0h-28c-29 0-51 2-63 5c-22 6-40 17-54 31c-13 14-24 32-31 56c-6 24-10 64-10 118c-1 55-1 91-2 105c-2 32-7 56-15 73c-8 16-21 30-37 40c-12 7-29 11-49 12v81c25 1 44 8 59 19c16 12 28 28 34 51c6 22 10 63 10 124c0 60 2 100 5 122c4 32 13 57 26 75c14 18 31 32 54 40c15 6 39 8 73 8h28v-77c-33 0-55-3-65-7s-17-9-21-15c-7-10-12-24-14-42c-1-11-2-49-2-114c0-72-8-122-24-154s-43-56-80-71c33-15 57-33 71-56c15-22 26-53 30-94c2-22 3-71 3-149c0-43 7-71 18-84c11-12 33-18 68-18h16V0z"/>`),
		g.Group(children),
	)
}