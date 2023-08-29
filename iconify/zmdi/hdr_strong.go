package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdrStrong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 64q53 0 90.5 37.5T469 192t-37.5 90.5T341 320t-90.5-37.5T213 192t37.5-90.5T341 64zM85.5 107q35.5 0 60.5 25t25 60t-25 60t-60.5 25T25 252T0 192t25-60t60.5-25zm0 128q17.5 0 30-12.5T128 192t-12.5-30.5t-30-12.5t-30 12.5T43 192t12.5 30.5t30 12.5z"/>`),
		g.Group(children),
	)
}