package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Renpho(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.835 42.5l18.169-17.688a12.445 12.445 0 0 0 9.938-3.36a10.589 10.589 0 0 0 3.187-9.417"/><ellipse cx="23.739" cy="11.52" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.184" ry="6.02"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.165 42.493L31.791 32.394c-7.795 3.081-15.218-.348-15.218-.348"/>`),
		g.Group(children),
	)
}