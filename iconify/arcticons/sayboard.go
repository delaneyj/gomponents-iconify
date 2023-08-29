package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sayboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.689 32.36l11.822-.067a2 2 0 0 0 1.989-2V8.94a2 2 0 0 0-2-2H7.626a2 2 0 0 0-2 1.988L5.5 30.283a2 2 0 0 0 1.987 2.012l9.316.066v8.7l11.886-8.7Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="3.806" height="4.486" x="8.784" y="10.145" rx="1" ry="1"/><rect width="3.806" height="4.486" x="15.438" y="10.145" rx="1" ry="1"/><rect width="3.806" height="4.486" x="22.091" y="10.145" rx="1" ry="1"/><rect width="3.806" height="4.486" x="28.745" y="10.145" rx="1" ry="1"/><rect width="3.806" height="4.486" x="35.399" y="10.145" rx="1" ry="1"/><rect width="3.806" height="4.486" x="8.784" y="23.912" rx="1" ry="1"/><rect width="17.113" height="4.486" x="15.438" y="23.912" rx="1" ry="1"/><rect width="3.806" height="4.486" x="35.399" y="23.912" rx="1" ry="1"/><rect width="3.806" height="4.486" x="12.111" y="17.029" rx="1" ry="1"/><rect width="3.806" height="4.486" x="18.764" y="17.029" rx="1" ry="1"/><rect width="3.806" height="4.486" x="25.418" y="17.029" rx="1" ry="1"/><rect width="3.806" height="4.486" x="32.072" y="17.029" rx="1" ry="1"/></g>`),
		g.Group(children),
	)
}