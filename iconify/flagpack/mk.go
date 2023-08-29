package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackMk0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackMk2)"><use href="#flagpackMk0"/><path fill="#F50100" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackMk1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill="#FFD018" mask="url(#flagpackMk1)"><path fill-rule="evenodd" d="M0-.042v4.084l14 5.951L3.255-.042H0ZM16 12l3-12h-6l3 12Zm0 0l-3 12h6l-3-12ZM0 19.951v4.084h3.255L14 14L0 19.951ZM32 4.066V-.018h-3.255L18 10.018l14-5.952Zm0 19.994v-4.084l-14-5.952L28.745 24.06H32ZM32 9l-12 3l12 3V9Zm-20 3L0 9v6l12-3Z" clip-rule="evenodd"/><path stroke="#F50100" stroke-width="2" d="M16 17a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/></g></g><defs><clipPath id="flagpackMk2"><use href="#flagpackMk0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}