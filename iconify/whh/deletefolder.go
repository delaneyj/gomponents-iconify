package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deletefolder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M988.94 855q9 35-15 60l-90 89q-25 25-61 16q-36 9-61-16l-89-89l-90 89q-25 25-61 16q-36 9-61-16l-90-89q-7-8-12-19h-230q-53 0-90.5-37.5T.94 768V256q0-27 19-45.5t45-18.5h480l46-85q9-16 30-29.5t42-13.5h240q24 0 45 12.5t28 30.5q49 68 49 85v576q0 49-36 87zm-110-239q18-18 18-43.5t-18-43.5t-43.5-18t-43.5 18l-119 119l-120-119q-18-18-43.5-18t-43.5 18t-18 43.5t18 43.5l119 119l-119 119q-18 18-18 43.5t18 43.5t43.5 18t43.5-18l120-119l119 119q18 18 43.5 18t43.5-18t18-43.5t-18-43.5l-119-119zm-832-573q7-18 28-30.5t44-12.5h240q24 0 45 12.5t29 30.5l48 85H.94z"/>`),
		g.Group(children),
	)
}