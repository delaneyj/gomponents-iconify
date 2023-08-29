package skill_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngularLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="none"><rect width="256" height="256" fill="#F4F2ED" rx="60"/><path fill="#E23237" d="M34.25 61.125L127.325 28l95.525 32.612L207.412 183.7L127.325 228l-78.787-43.725L34.25 61.125Z"/><path fill="#B52E31" d="M222.85 60.612L127.325 28v200l80.087-44.3L222.85 60.613Z"/><path fill="#fff" d="m127.469 51.375l-58 129l21.644-.438l11.687-29.149h51.875l12.7 29.375l20.637.437l-60.543-129.225Zm.143 41.412l19.625 40.982H110.5l17.169-40.982h-.057Z"/></g>`),
		g.Group(children),
	)
}