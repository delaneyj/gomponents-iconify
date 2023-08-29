package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlepause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm-64-704q0-26-18.5-45T384 256h-64q-26 0-45 19t-19 45v384q0 27 19 45.5t45 18.5h64q27 0 45.5-18.5T448 704V320zm320 0q0-26-18.5-45T704 256h-64q-26 0-45 19t-19 45v384q0 27 19 45.5t45 18.5h64q27 0 45.5-18.5T768 704V320z"/>`),
		g.Group(children),
	)
}