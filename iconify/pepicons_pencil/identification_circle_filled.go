package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdentificationCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilIdentificationCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M7 8a1.5 1.5 0 0 0-1.5 1.5v7A1.5 1.5 0 0 0 7 18h12a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 19 8H7ZM4.5 9.5A2.5 2.5 0 0 1 7 7h12a2.5 2.5 0 0 1 2.5 2.5v7A2.5 2.5 0 0 1 19 19H7a2.5 2.5 0 0 1-2.5-2.5v-7Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.182 10.818a.5.5 0 0 1 .5-.5h2.727a.5.5 0 0 1 0 1h-2.727a.5.5 0 0 1-.5-.5Zm0 4.092a.5.5 0 0 1 .5-.5h5.454a.5.5 0 1 1 0 1h-5.454a.5.5 0 0 1-.5-.5Zm0-2.728a.5.5 0 0 1 .5-.5h5.454a.5.5 0 0 1 0 1h-5.454a.5.5 0 0 1-.5-.5Zm0 1.364a.5.5 0 0 1 .5-.5h5.454a.5.5 0 1 1 0 1h-5.454a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path d="M10.773 12.182a1.364 1.364 0 1 1-2.728 0a1.364 1.364 0 0 1 2.728 0Z"/><path d="M11.045 14.775c0 .688-.732.623-1.636.623c-.904 0-1.636.066-1.636-.623c0-.688.732-1.557 1.636-1.557c.904 0 1.636.87 1.636 1.557Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilIdentificationCircleFilled0)"/></g>`),
		g.Group(children),
	)
}