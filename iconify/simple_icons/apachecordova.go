package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apachecordova(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.545.545H5.455L0 9.273l2.182 14.182h3.886l-.273-3.273h1.909l.273 3.273h8.045l.273-3.273h1.909l-.273 3.273h3.886L24 9.273L18.545.545zm0 17.455H5.455L4.364 9.273l2.182-4.364h3.506l-.234 1.636h4.364l-.234-1.636h3.506l2.182 4.364L18.545 18zm-3-6.955c.301 0 .545.908.545 2.029s-.244 2.029-.545 2.029c-.301 0-.545-.908-.545-2.029c0-1.12.244-2.029.545-2.029zm-6.886.17c.301 0 .545.908.545 2.029s-.244 2.029-.545 2.029c-.301 0-.545-.908-.545-2.029s.244-2.029.545-2.029z"/>`),
		g.Group(children),
	)
}