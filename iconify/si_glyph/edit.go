package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Edit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="m15.682 2.91l-1.535-1.514c-.41-.408-1.218-.577-1.607-.189L6.485 7.234s-1.321 3.913-1.446 4.242c-.099.262.231.59.46.475c.374-.186 4.105-1.631 4.105-1.631l6.071-6.057c.39-.388.418-.946.007-1.353zM9.021 9.854l-2.567.974l-.461-.425l1.225-2.965l.741.546h1.058l.004 1.87z"/><path d="M12.042 10.345v3.697H.958V3l8.183-.083l.812-.859H0V15h13l-.057-5.422l-.901.767z"/></g>`),
		g.Group(children),
	)
}