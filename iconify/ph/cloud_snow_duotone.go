package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSnowDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M224 92a68 68 0 0 1-68 68H76a44 44 0 1 1 14.2-85.66v.11A68.06 68.06 0 0 1 224 92Z" opacity=".2"/><path d="M88 196a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm28 4a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm48-16a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm-96 40a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm88 0a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm76-132a76.08 76.08 0 0 1-76 76H76a52 52 0 0 1 0-104a53.26 53.26 0 0 1 8.92.76A76.08 76.08 0 0 1 232 92Zm-16 0a60.06 60.06 0 0 0-120-3.54a8 8 0 0 1-16-.92q.21-3.66.77-7.23A38.11 38.11 0 0 0 76 80a36 36 0 0 0 0 72h80a60.07 60.07 0 0 0 60-60Z"/></g>`),
		g.Group(children),
	)
}