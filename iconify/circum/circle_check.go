package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.81 10.4a.5.5 0 0 0-.71-.71l-3.56 3.56l-1.73-1.73a.5.5 0 0 0-.71.71l2.08 2.08a.513.513 0 0 0 .71 0Z"/><path fill="currentColor" d="M12 21.934A9.934 9.934 0 1 1 21.933 12A9.945 9.945 0 0 1 12 21.934Zm0-18.867A8.934 8.934 0 1 0 20.933 12A8.944 8.944 0 0 0 12 3.067Z"/>`),
		g.Group(children),
	)
}