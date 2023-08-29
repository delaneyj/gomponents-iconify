package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommandDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 76a28 28 0 0 1-28 28h-28V76a28 28 0 0 1 28-28a28 28 0 0 1 28 28ZM76 48a28 28 0 0 0-28 28a28 28 0 0 0 28 28h28V76a28 28 0 0 0-28-28Zm104 104h-28v28a28 28 0 0 0 28 28a28 28 0 0 0 28-28a28 28 0 0 0-28-28ZM48 180a28 28 0 0 0 28 28a28 28 0 0 0 28-28v-28H76a28 28 0 0 0-28 28Z" opacity=".2"/><path d="M180 144h-20v-32h20a36 36 0 1 0-36-36v20h-32V76a36 36 0 1 0-36 36h20v32H76a36 36 0 1 0 36 36v-20h32v20a36 36 0 1 0 36-36Zm-20-68a20 20 0 1 1 20 20h-20ZM56 76a20 20 0 0 1 40 0v20H76a20 20 0 0 1-20-20Zm40 104a20 20 0 1 1-20-20h20Zm16-68h32v32h-32Zm68 88a20 20 0 0 1-20-20v-20h20a20 20 0 0 1 0 40Z"/></g>`),
		g.Group(children),
	)
}