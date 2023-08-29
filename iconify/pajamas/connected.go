package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Connected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.78 2.28a.75.75 0 0 0-1.06-1.06l-1.617 1.616a4 4 0 0 0-5.275.336L5.53 4.47l-.25-.25a.75.75 0 0 0-1.06 1.06l.25.25l-1.302 1.302a4 4 0 0 0-.336 5.275L1.22 13.72a.75.75 0 1 0 1.06 1.06l1.613-1.612a4 4 0 0 0 5.275-.336l1.302-1.302l.25.25a.75.75 0 1 0 1.06-1.06l-.25-.25l1.298-1.298a4 4 0 0 0 .336-5.275L14.78 2.28Zm-4.31 7.13l1.297-1.297a2.5 2.5 0 0 0 0-3.536l-.343-.343l.404-.405l-.404.405a2.5 2.5 0 0 0-3.536 0L6.591 5.53l3.879 3.88ZM5.53 6.591l-1.3 1.301a2.5 2.5 0 0 0 0 3.536l-.4.4l.4-.4l.343.343a2.5 2.5 0 0 0 3.536 0L9.41 10.47L5.53 6.59Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}