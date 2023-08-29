package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dizzy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="m59.104 24.766l-13.6 3.182l-10.344-9.41l-1.172 13.941l-12.113 6.928l7.971 3.36c-8.95 9.691-21.285 16.05-22.233 12.592c-1.375-5.025 17.34-32.826 35.484-44.744C60.908-1.083 61.4 8.561 59.682 13.795C68.557-5.248 49.547 3.598 43.063 7.336C19.68 20.809-3.102 54.676 3.004 60.602c4.405 4.276 22.777-2.003 33.211-8.733l1.389 6.671l9.119-10.584l13.891 1.535l-7.244-11.977l5.734-12.748"/>`),
		g.Group(children),
	)
}