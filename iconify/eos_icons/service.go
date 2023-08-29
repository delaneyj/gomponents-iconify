package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Service(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21h2v-6H3v6zm9.46 1l-2.93-.64l-2.93-.69a.73.73 0 0 1-.43-.26a.71.71 0 0 1-.17-.47v-4a.77.77 0 0 1 .17-.47a.8.8 0 0 1 .43-.26l4.78-.59l4.79-.6l.15.56l.15.55a.93.93 0 0 1-.15.52a.9.9 0 0 1-.4.37l-2 .5l-2 .5l1.62.66l1.63.65l2.27-.64l2.21-.69a.78.78 0 0 1 .58.08a.73.73 0 0 1 .36.46l.19.77l.2.76a.77.77 0 0 1-.08.56a.78.78 0 0 1-.46.35l-3.63 1l-3.63 1a1.42 1.42 0 0 1-.36 0a1.1 1.1 0 0 1-.36.02ZM9 4L7 6L5 8l2 2l2 2l.7-.7l.7-.7l-1.3-1.3L7.8 8l1.3-1.3l1.3-1.3l-.7-.7L9 4zm6 0l-.7.7l-.7.7l1.3 1.3L16.2 8l-1.3 1.3l-1.3 1.3l.7.7l.7.7l2-2l2-2l-2-2l-2-2z"/>`),
		g.Group(children),
	)
}