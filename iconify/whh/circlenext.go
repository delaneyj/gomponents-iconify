package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlenext(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm-64-704q0-26-18.5-45T384 256h-64q-27 0-45.5 19T256 320v384q0 27 18.5 45.5T320 768h64q27 0 45.5-18.5T448 704V320zm377 173L607 262q-13-14-31 6v488q18 20 31 7l218-232q7-7 7-18.5t-7-19.5z"/>`),
		g.Group(children),
	)
}