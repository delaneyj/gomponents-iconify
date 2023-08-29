package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Batman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768.424 896q-49 49-154 88.5t-194 39.5q-91 0-199-39.5t-157-88.5q-64-64-64-512V32q0-13 9.5-22.5t22.5-9.5t22.5 9.5t9.5 22.5l37 149q60-54 143.5-85.5t175.5-31.5q91 0 172 31t139 84l37-147q0-13 9.5-22.5t22.5-9.5t22.5 9.5t9.5 22.5v352q0 448-64 512zm-512-448q-26 0-65.5-20t-53.5-46q-9 17-9 34q0 43 23 69.5t73 26.5q48 0 68.5-14t27.5-50h-64zm439-66q-14 27-52.5 46.5t-66.5 19.5h-64q7 36 27.5 50t68.5 14q50 0 73-26.5t23-69.5q0-17-9-34zm-279 322l-352-128q0 218 54 273q41 42 133 76.5t168 34.5q75 0 164-34t131-77q54-55 54-273z"/>`),
		g.Group(children),
	)
}