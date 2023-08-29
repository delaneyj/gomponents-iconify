package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleExclamationMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="#d22f27"><ellipse cx="26.669" cy="59.435" rx="3" ry="3.023"/><path d="M26.57 49.964a2.485 2.485 0 0 1-2.5-2.519V11.977a2.5 2.5 0 1 1 5 0v35.468a2.485 2.485 0 0 1-2.5 2.519Z"/><ellipse cx="45.331" cy="59.435" rx="3" ry="3.023"/><path d="M45.23 49.964a2.485 2.485 0 0 1-2.5-2.519V11.977a2.5 2.5 0 1 1 5 0v35.468a2.485 2.485 0 0 1-2.5 2.519Z"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="26.669" cy="60.081" rx="3" ry="3.023"/><ellipse cx="45.331" cy="60.081" rx="3" ry="3.023"/><path d="M26.67 49.964a2.485 2.485 0 0 1-2.5-2.519V11.977a2.5 2.5 0 1 1 5 0v35.468a2.485 2.485 0 0 1-2.5 2.519Zm18.66 0a2.485 2.485 0 0 1-2.5-2.519V11.977a2.5 2.5 0 1 1 5 0v35.468a2.485 2.485 0 0 1-2.5 2.519Z"/></g>`),
		g.Group(children),
	)
}