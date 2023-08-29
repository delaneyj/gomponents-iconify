package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stroller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 512H684L504 691l118 141h99q17-29 46.5-46.5T832 768q53 0 90.5 37.5T960 896t-37.5 90.5T832 1024t-90.5-37.5T704 896h-96q-18 0-27-16L460 736l-95 95q19 31 19 65q0 53-37.5 90.5T256 1024t-90.5-37.5T128 896t37.5-90.5T256 768q34 0 65 18l99-99l-146-175h-82l-33-117L9 244q-9-9-9-21.5T9 201t21.5-9t21.5 9l119 119h405L512 0q164 0 313 78t199 178zm-64.5 448q26.5 0 45.5-19t19-45.5t-19-45t-45.5-18.5t-45 18.5t-18.5 45t18.5 45.5t45 19zm-576-128q-26.5 0-45 18.5t-18.5 45t18.5 45.5t45 19t45.5-19t19-45.5t-19-45t-45.5-18.5zM355 512l109 131l132-131H355z"/>`),
		g.Group(children),
	)
}