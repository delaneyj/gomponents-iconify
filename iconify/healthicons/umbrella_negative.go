package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsUmbrellaNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24.976 5a1 1 0 1 0-2 0v1.029c-9.144.5-16.47 7.826-16.97 16.97c-.017.293.339.343.557.147a7.172 7.172 0 0 1 4.815-1.847c2.17 0 4.116.96 5.436 2.48c.119.137.289.22.47.22c.16 0 .313-.065.428-.177a8.97 8.97 0 0 1 5.264-2.47V39.5a4.5 4.5 0 1 0 9 0V38a1 1 0 1 0-2 0v1.5a2.5 2.5 0 0 1-5 0V21.356a8.969 8.969 0 0 1 5.223 2.456a.655.655 0 0 0 .453.187c.191 0 .37-.088.496-.232a7.182 7.182 0 0 1 5.426-2.468c1.851 0 3.54.699 4.815 1.847c.218.196.573.146.557-.147c-.5-9.144-7.826-16.47-16.97-16.97V5Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsUmbrellaNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}