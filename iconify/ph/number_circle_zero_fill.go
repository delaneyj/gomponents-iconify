package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleZeroFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M156 128c0 14.86-5.9 40-28 40s-28-25.14-28-40s5.9-40 28-40s28 25.14 28 40Zm76 0A104 104 0 1 1 128 24a104.11 104.11 0 0 1 104 104Zm-60 0c0-14.25-3.56-27.53-10-37.39C154 78.44 142.23 72 128 72s-26 6.44-34 18.61c-6.47 9.86-10 23.14-10 37.39s3.56 27.53 10 37.39c8 12.18 19.74 18.61 34 18.61s26-6.43 34-18.61c6.44-9.86 10-23.14 10-37.39Z"/>`),
		g.Group(children),
	)
}