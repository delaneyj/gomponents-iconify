package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M890.585 221q6 20 6 36v512h64q26 0 45 18.5t19 45t-18.5 45.5t-45.5 19h-64v64q0 27-19 45.5t-45.5 18.5t-45-18.5t-18.5-45.5v-64h-512q-53 0-90.5-37.5t-37.5-90.5V257h-64q-26 0-45-18.5t-19-45.5t18.5-45.5t45.5-18.5h64V65q0-26 18.5-45t45.5-19t45.5 19t18.5 45v64h512q16 0 35 5l116-116q18-18 43.5-18t43.5 18t18 43.5t-18 43.5zm-122 122l-426 426h426V343zm-512-86v424l424-424h-424z"/>`),
		g.Group(children),
	)
}