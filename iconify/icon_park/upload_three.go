package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M39 6H9C7.34315 6 6 7.34315 6 9V39C6 40.6569 7.34315 42 9 42H39C40.6569 42 42 40.6569 42 39V9C42 7.34315 40.6569 6 39 6Z"/><path stroke="#fff" stroke-linecap="round" d="M32 28L24 20L16 28"/><path stroke="#fff" stroke-linecap="round" d="M24 36V20.5"/><path stroke="#fff" stroke-linecap="round" d="M16 14H32"/></g>`),
		g.Group(children),
	)
}