package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendCharts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M128 896V128h768v768H128zm291.712-327.296l128 102.4l180.16-201.792l-47.744-42.624l-139.84 156.608l-128-102.4l-180.16 201.792l47.744 42.624l139.84-156.608zM816 352a48 48 0 1 0-96 0a48 48 0 0 0 96 0z"/>`),
		g.Group(children),
	)
}