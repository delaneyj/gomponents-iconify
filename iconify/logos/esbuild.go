package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Esbuild(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<circle cx="128" cy="128" r="128" fill="#FFCF00"/><path fill="#191919" d="M69.285 58.715L138.571 128l-69.286 69.285l-16.97-16.97L104.629 128L52.315 75.685l16.97-16.97Zm76.8 0L215.371 128l-69.286 69.285l-16.97-16.97L181.429 128l-52.314-52.315l16.97-16.97Z"/>`),
		g.Group(children),
	)
}