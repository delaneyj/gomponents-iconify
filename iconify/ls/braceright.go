package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Braceright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M289 521v-81c-25-1-45-7-61-19c-15-11-27-28-33-50c-7-22-10-64-10-125c-1-59-2-101-5-123c-4-31-13-56-27-74c-13-19-30-33-53-40c-15-6-40-9-73-9H0v79h15c30 0 50 3 60 10s17 20 21 37c3 12 4 55 4 129c0 71 9 120 24 151s42 56 81 74c-29 12-53 30-69 53s-27 55-32 95c-2 22-4 72-4 152c0 43-5 71-16 83s-34 20-69 20H0v77h27c29 0 50-1 63-5c22-6 40-16 53-30c14-13 25-33 32-57c6-24 10-63 10-118c1-55 1-90 2-105c1-32 7-56 16-73c8-16 20-30 36-40c12-7 30-11 50-11z"/>`),
		g.Group(children),
	)
}