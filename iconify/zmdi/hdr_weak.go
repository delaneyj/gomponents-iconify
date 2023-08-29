package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdrWeak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M85.5 107q35.5 0 60.5 25t25 60t-25 60t-60.5 25T25 252T0 192t25-60t60.5-25zM341 64q53 0 90.5 37.5T469 192t-37.5 90.5T341 320t-90.5-37.5T213 192t37.5-90.5T341 64zm.5 213q35.5 0 60.5-25t25-60t-25-60t-60.5-25t-60.5 25t-25 60t25 60t60.5 25z"/>`),
		g.Group(children),
	)
}