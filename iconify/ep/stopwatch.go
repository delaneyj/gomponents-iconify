package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stopwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 896a384 384 0 1 0 0-768a384 384 0 0 0 0 768zm0 64a448 448 0 1 1 0-896a448 448 0 0 1 0 896z"/><path fill="currentColor" d="M672 234.88c-39.168 174.464-80 298.624-122.688 372.48c-64 110.848-202.624 30.848-138.624-80C453.376 453.44 540.48 355.968 672 234.816z"/>`),
		g.Group(children),
	)
}