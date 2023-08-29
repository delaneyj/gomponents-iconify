package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePixelBuds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="arcticonsGooglePixelBuds0" fill="none" stroke="currentColor" stroke-miterlimit="10" d="M13.838 39.917a2.785 2.785 0 0 1-1.898-4.825c.099-.095 4.829-4.648 4.829-11.092c0-6.557-4.773-11.04-4.822-11.085a2.786 2.786 0 0 1 3.78-4.093c.27.25 6.613 6.21 6.613 15.178s-6.343 14.929-6.613 15.178a2.78 2.78 0 0 1-1.89.739Z"/></defs><use href="#arcticonsGooglePixelBuds0" stroke-miterlimit="10"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M7.288 33.168a2.786 2.786 0 0 1-1.714-4.984c.123-.1 1.885-1.587 1.885-4.184c0-2.634-1.812-4.126-1.889-4.188c-1.185-.957-1.4-2.701-.46-3.899a2.769 2.769 0 0 1 3.861-.514c1.508 1.146 4.06 4.247 4.06 8.6s-2.552 7.456-4.06 8.601a2.772 2.772 0 0 1-1.683.568Z"/><use href="#arcticonsGooglePixelBuds0" stroke-miterlimit="10"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M7.288 33.168a2.786 2.786 0 0 1-1.714-4.984c.123-.1 1.885-1.587 1.885-4.184c0-2.634-1.812-4.126-1.889-4.188c-1.185-.957-1.4-2.701-.46-3.899a2.769 2.769 0 0 1 3.861-.514c1.508 1.146 4.06 4.247 4.06 8.6s-2.552 7.456-4.06 8.601a2.772 2.772 0 0 1-1.683.568Zm26.874 6.749a2.785 2.785 0 0 0 1.898-4.825c-.1-.095-4.83-4.648-4.83-11.092c0-6.557 4.773-11.04 4.822-11.085a2.786 2.786 0 0 0-3.78-4.093c-.27.25-6.613 6.21-6.613 15.178s6.343 14.929 6.613 15.178a2.78 2.78 0 0 0 1.89.739Z"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M40.712 33.168a2.786 2.786 0 0 0 1.714-4.984c-.123-.1-1.885-1.587-1.885-4.184c0-2.634 1.812-4.126 1.889-4.188c1.185-.957 1.4-2.701.46-3.899a2.769 2.769 0 0 0-3.861-.514c-1.508 1.146-4.06 4.247-4.06 8.6s2.552 7.456 4.06 8.601a2.772 2.772 0 0 0 1.683.568Z"/>`),
		g.Group(children),
	)
}