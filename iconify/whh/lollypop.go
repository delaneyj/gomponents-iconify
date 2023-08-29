package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lollypop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1007 1007q-17 17-41 17t-41-17L502 584q-82 56-181 56q-24 0-53-4l368-368q4 29 4 53q0 99-56 181l423 423q17 17 17 41t-17 41zM0 465L465 0q59 26 104.5 71.5T641 176L176 641q-59-26-104.5-71.5T0 465zm1-144q0-87 42.5-161T160 43.5T321 1q24 0 52 4L5 373q-4-28-4-52z"/>`),
		g.Group(children),
	)
}