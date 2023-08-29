package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PercentDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M95.8 56.2a28 28 0 1 1-39.6 0a28 28 0 0 1 39.6 0Zm104 104a28 28 0 1 0 0 39.6a28 28 0 0 0 0-39.6Z" opacity=".2"/><path d="m205.66 61.64l-144 144a8 8 0 0 1-11.32-11.32l144-144a8 8 0 0 1 11.32 11.31Zm-155.12 39.8a36 36 0 0 1 50.92-50.91a36 36 0 0 1-50.92 50.91ZM56 76a20 20 0 1 0 34.14-14.16A20 20 0 0 0 56 76Zm160 104a36 36 0 1 1-10.54-25.46A35.76 35.76 0 0 1 216 180Zm-16 0a20 20 0 1 0-5.86 14.14A19.87 19.87 0 0 0 200 180Z"/></g>`),
		g.Group(children),
	)
}