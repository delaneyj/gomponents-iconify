package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneEnabledOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.05 21q-.45 0-.75-.3t-.3-.75V15.9q0-.325.225-.588t.575-.362l3.45-.7q.35-.05.713.063t.587.337L10.9 17q.95-.55 1.8-1.213t1.625-1.437q.825-.8 1.513-1.663t1.187-1.787L14.6 8.45q-.2-.2-.275-.475T14.3 7.3l.65-3.5q.05-.325.325-.563T15.9 3h4.05q.45 0 .75.3t.3.75q0 3.125-1.363 6.175t-3.862 5.55q-2.5 2.5-5.55 3.863T4.05 21Zm13.9-12q.425-.975.65-1.975T18.95 5h-2.2l-.45 2.35L17.95 9ZM9 17.9l-1.65-1.65l-2.35.5v2.2q1.025-.075 2.025-.35T9 17.9ZM17.95 9ZM9 17.9Z"/>`),
		g.Group(children),
	)
}