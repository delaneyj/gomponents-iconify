package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m955.48 852l-103 103q-69 69-167 69t-167.5-69t-69.5-167t69-167l38-38l-171-172l-83 83q-18 18-44 17t-45-20t-20-45t17-44l83-82l-101-101l-82 83q-18 18-44 17t-45-20t-20-45t17-44l82-82l-19-19q-17-18-16.5-44t19.5-45t45.5-20t43.5 17l478 477q69-51 156-44.5t149 68.5q69 69 69 167t-69 167zm-91-243q-33-33-79.5-33t-79.5 33l-96 96q-33 33-33 79.5t33 79.5t80 33t80-33l95-95q33-33 33-80t-33-80z"/>`),
		g.Group(children),
	)
}