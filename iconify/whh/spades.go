package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spades(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 896q-112 0-188-83q7 94 23.5 152.5T640 1024H384q20 0 36.5-58.5T444 813q-76 83-188 83q-106 0-181-75T0 640q0-78 28-149t73.5-124t100-99.5t109-83.5t100-68T484 55t28-55q0 25 28 55t73.5 61t100 68t109 83.5t100 99.5T996 491t28 149q0 106-75 181t-181 75z"/>`),
		g.Group(children),
	)
}