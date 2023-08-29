package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyMinimalisticSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12.39 9.724a1.333 1.333 0 1 1 1.886 1.885a1.333 1.333 0 0 1-1.886-1.885Z"/><path fill-rule="evenodd" d="M3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464ZM16.651 7.35a4.605 4.605 0 0 1-4.405 7.715a.72.72 0 0 0-.691.172l-1.118 1.118l-.003-.003l-1.166-1.153a.75.75 0 1 0-1.055 1.066l1.163 1.15l-.296.297a.987.987 0 0 1-.807.283l-1.392-.155a.657.657 0 0 1-.393-.188l-.139-.14a.658.658 0 0 1-.188-.392l-.155-1.392a.987.987 0 0 1 .283-.807l2.475-2.474a.72.72 0 0 0 .172-.692a4.604 4.604 0 0 1 7.715-4.405Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}