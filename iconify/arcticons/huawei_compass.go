package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuaweiCompass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.351 5.887c9.586 3.372 18.343 6.314 29.963 10.414c3.31 1.168 2.612 3.828-.427 4.202c-5.675.699-8.982 1.233-12.423 1.705c-3.328.457-6.25 3.838-7.186 7.065c-.906 3.123-1.693 6.679-2.619 10.718c-.593 2.587-3.411 3.566-4.263.183c-2.31-9.184-5.095-20.392-7.734-29.963a3.506 3.506 0 0 1 4.69-4.323Z"/>`),
		g.Group(children),
	)
}