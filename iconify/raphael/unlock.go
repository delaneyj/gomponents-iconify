package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.875 15.334v-4.876c0-4.894-3.98-8.875-8.875-8.875s-8.875 3.98-8.875 8.875v.375h3.5v-.375c0-2.964 2.41-5.375 5.375-5.375s5.375 2.41 5.375 5.375v4.876H5.042v15.083h21.916V15.334h-2.083zm-6.603 11.622h-4.545l1.222-3.667a2.37 2.37 0 0 1-1.325-2.12a2.375 2.375 0 1 1 4.75 0c0 .932-.542 1.73-1.324 2.12l1.222 3.666z"/>`),
		g.Group(children),
	)
}