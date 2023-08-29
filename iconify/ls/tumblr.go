package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tumblr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M84 14h498c47 0 84 37 84 84v498c0 25-10 44-25 59c-14 14-36 25-59 25H84c-47 0-84-37-84-84V98c0-23 11-45 25-59c15-15 34-25 59-25zm141 275v186c0 25 2 43 7 55s15 25 28 36s29 21 48 27c18 6 40 9 64 9c22 0 41-2 59-7c18-4 39-12 62-22v-84c-27 17-53 27-80 27c-15 0-29-5-40-12c-9-5-16-11-19-20c-3-8-5-28-5-59V289h126v-82H349V72h-75c-3 27-9 50-18 67c-9 18-21 33-36 46c-14 12-31 22-53 29v75h58z"/>`),
		g.Group(children),
	)
}