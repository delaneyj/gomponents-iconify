package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.933A9.933 9.933 0 1 1 21.933 12A9.944 9.944 0 0 1 12 21.933Zm0-18.866A8.933 8.933 0 1 0 20.933 12A8.943 8.943 0 0 0 12 3.067Z"/><path fill="currentColor" d="M11.5 6a.5.5 0 0 1 1 0v4.8c1.13-1.13 2.26-2.27 3.39-3.4a.5.5 0 0 1 .71.71l-4.26 4.25a.463.463 0 0 1-.58.07c-.01-.02-.02-.02-.03-.02a.425.425 0 0 1-.22-.33Z"/>`),
		g.Group(children),
	)
}