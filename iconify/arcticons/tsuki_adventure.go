package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TsukiAdventure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.262 5.5L5.5 11.726m37 26.387L34.361 42.5m8.139-7.406l-13.924 7.381m1.598-4.426c3.125.873 6.691-2.915 4.32-6.065"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.361 34.315c-.799 3.586-4.597 4.929-6.716 5.007"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.975 29.628c.656 1.503.096 5.82-3.518 4.112"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.777 32.195l4.326-3.737c1.492.059 2.813-1.002 1.543-1.751c1.223-2.36 3.502-3.905 5.588-5.598m.107-2.942c-.618 1.203-1.797 2.171-2.776 2.938c-3.372 2.495-6.356 6.337-9.593 9.576"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.795 35.23c4.92-4.159-.995-6.546-3.699-4.241m5.849-4.973c.372-.044.58-.014.883-.367m-6.896-1.427l.256 1.2M12.71 35.127c-4.252.76-3.592-4.19-.841-3.697"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.709 39.924c-2.514-.245-8.152-2.994-6.461-13.031"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.636 34.92c-.712-.806-3.521-1.049-5.36 2.092c-1.431 2.446.533 6.582 5.13 3.508c1.17-.784 2.047-3.543.23-5.6ZM5.5 21.48c2.544.895 5.58 2.263 8.758.343m-1.333 4.447C9.896 29.23 5.8 28.211 5.5 28.067"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.608 30.63c-7.926 1.889-11.633-1.765-13.113-5.257c.729-2.55 1.294-2.773 2.132-3.98m5.578-7.554c7.818-4.402 16.085 7.075 12.649 13.763c-1.515 1.488-3.192 2.257-4.931 2.499"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.793 30.186c-2.072.03-4.202-.613-6.228-1.614c-4.24-1.297-7.255-6.764-5.204-10.878m7.682-4.885c-3.097-2.943-1.587-5.737.3-6.136c3.202-.678 6.833 7.626 6.617 8.58M14.007 12.48c2.566 1.375 4.06 3.54 5.038 5.288"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.197 19.284c-.442-1.227-8.027-5.407-4.326-8.933c2.882-2.745 8.322 4.868 8.96 7.022"/>`),
		g.Group(children),
	)
}