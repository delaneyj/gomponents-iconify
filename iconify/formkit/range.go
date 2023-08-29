package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Range(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2.31 2.5H.5C.22 2.5 0 2.28 0 2s.22-.5.5-.5h1.81c-.04.16-.07.33-.07.5s.03.34.07.5zm11.19-1H6.17c.04.16.07.33.07.5s-.03.34-.07.5h7.33c.28 0 .5-.22.5-.5s-.22-.5-.5-.5zm-7.33 1c-.22.86-1 1.5-1.93 1.5s-1.71-.64-1.93-1.5c-.04-.16-.07-.33-.07-.5s.03-.34.07-.5C2.53.64 3.31 0 4.24 0s1.71.64 1.93 1.5c.04.16.07.33.07.5s-.03.34-.07.5zM5.24 2c0-.55-.45-1-1-1s-1 .45-1 1s.45 1 1 1s1-.45 1-1z"/>`),
		g.Group(children),
	)
}