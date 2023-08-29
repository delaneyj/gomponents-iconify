package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OperaAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 5.5c-3.1 0-3.1 4.7-3.1 6.4c0 1.8 0 6.6 3.2 6.6s3.2-4.8 3.2-6.6c-.2-4.2-1.2-6.4-3.3-6.4zm0 12c-1.5 0-2.2-1.8-2.2-5.7c0-3.6.7-5.3 2.1-5.3c1.4 0 2.2 1.8 2.2 5.4c0 3.8-.6 5.6-2.1 5.6zm0-16c-5.8 0-9.8 4.3-9.8 10.4c0 5.2 3.7 10.6 9.8 10.6c6.1 0 9.8-5.4 9.8-10.6c0-6.1-4-10.4-9.8-10.4zm0 20c-5.5 0-8.8-4.9-8.8-9.6c0-5.6 3.5-9.4 8.8-9.4c5.3 0 8.8 3.8 8.8 9.4c0 4.7-3.3 9.6-8.8 9.6z"/>`),
		g.Group(children),
	)
}