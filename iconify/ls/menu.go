package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Menu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M44 128h578c25 0 44-19 44-44c0-24-19-43-44-43H44C19 41 0 60 0 84c0 25 19 44 44 44zm0 175h578c25 0 44-19 44-44c0-24-19-43-44-43H44c-25 0-44 19-44 43c0 25 19 44 44 44zm0 175h578c25 0 44-19 44-43c0-25-19-44-44-44H44c-25 0-44 19-44 44c0 24 19 43 44 43zm0 175h578c25 0 44-19 44-43c0-25-19-44-44-44H44c-25 0-44 19-44 44c0 24 19 43 44 43z"/>`),
		g.Group(children),
	)
}