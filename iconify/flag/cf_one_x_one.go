package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CfOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagCf1x10"><path fill-opacity=".7" d="M0 0h512v512H0z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagCf1x10)"><path fill="#00f" d="M-52-.5h768v127H-52z"/><path fill="#ff0" d="M-52 383.5h768V512H-52z"/><path fill="#009a00" d="M-52 255h768v128.5H-52z"/><path fill="#fff" d="M-52 126.5h768V255H-52z"/><path fill="red" d="M268 0h128v512H268z"/><path fill="#ff0" d="M109.5 112.3L75.9 89.1l-33.4 23.4l11.6-39.2l-32.5-24.6l40.7-1L75.7 8.8l13.5 38.6l40.8.8L97.6 73"/></g>`),
		g.Group(children),
	)
}