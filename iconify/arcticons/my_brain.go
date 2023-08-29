package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyBrain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.682 19.796l2.803-11.213L34.09 5.78l5.606 5.606l-5.606 5.606ZM34.09 5.78l-8.408 14.015"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.697 11.386L42.5 22.6v8.409l-2.803 8.409l-5.606 2.803l-8.41-5.606V19.796l5.607 8.409l2.803-11.213L42.5 22.6l-11.212 5.606l-5.606 8.409"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.682 36.614h8.409l8.409-5.606"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.894 25.401L34.09 36.614l-2.803-8.41m-8.969-8.408L19.515 8.583L13.91 5.78l-5.606 5.606l5.606 5.606ZM13.91 5.78l8.408 14.015"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.303 11.386L5.5 22.6v8.409l2.803 8.409l5.606 2.803l8.41-5.606V19.796l-5.607 8.409l-2.803-11.213L5.5 22.6l11.212 5.606l5.606 8.409"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.318 36.614H13.91L5.5 31.008"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.106 25.401l2.803 11.213l2.803-8.41"/>`),
		g.Group(children),
	)
}