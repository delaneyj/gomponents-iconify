package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapPound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M52 2H12C6.478 2 2 6.477 2 11.999V52c0 5.522 4.478 10 10 10h40c5.522 0 10-4.478 10-10V11.999C62 6.477 57.522 2 52 2zm5 43.666A8.333 8.333 0 0 1 48.667 54H15.333A8.333 8.333 0 0 1 7 45.666V12.333A8.332 8.332 0 0 1 15.333 4h33.334A8.332 8.332 0 0 1 57 12.333v33.333z"/><path fill="currentColor" d="M49.365 24.921L51 18.941h-6.104l1.635-5.979H40.43l-1.637 5.979h-8.137l1.635-5.979h-6.102l-1.637 5.979H18.45l-1.635 5.979h6.102l-2.18 7.973h-6.104L13 38.873h6.104l-1.637 5.979h6.104l1.635-5.979h8.139l-1.637 5.979h6.104l1.635-5.979h6.104l1.635-5.979h-6.102l2.18-7.973h6.101m-14.386 7.973h-8.137l2.18-7.973h8.137l-2.18 7.973"/>`),
		g.Group(children),
	)
}