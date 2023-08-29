package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sky(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><circle cx="33.542" cy="21.309" r="2.809" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="19.427" cy="22.422" r="2.809" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.06 24.708a43.629 43.629 0 0 1 8.751 9.978a81.107 81.107 0 0 1-12.345.78c-3.048-.201-8.896-3.089-8.896-3.089c5.662-2.777 7.843-5.146 9.564-7.461"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.333 35.352l.563 2.199a.64.64 0 0 1-.609.8l-6.398.118a.618.618 0 0 1-.61-.77l.586-2.3"/><circle cx="24.712" cy="10.898" r="1.431" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.675 23.88v1.39s2.178 1.285 2.394 2.393a50.104 50.104 0 0 1 .122 9.483a1.614 1.614 0 0 1-1.595 1.467h-4.109a1.592 1.592 0 0 1-1.576-1.468a57.724 57.724 0 0 1 .103-9.324c.188-1.157 2.487-2.486 2.487-2.486v-1.417M42.5 35.505l-5.223-.005m-20.1.005L5.5 35.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.5 23.918c-.066 4.048 1.354 8.422.517 12.221l.88 2.474l2.073-2.08c.163-4.809-1.678-8.559-1.294-12.654m-13.616.829a19.581 19.581 0 0 1 .5 4.694m-3.427-4.486a19.576 19.576 0 0 1-3.815 2.78"/>`),
		g.Group(children),
	)
}