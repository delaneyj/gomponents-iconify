package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Questionsign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm.5-128q26.5 0 45-18.5t18.5-45t-18.5-45.5t-45-19t-45.5 19t-19 45.5t19 45t45.5 18.5zm-.5-768q-106 0-181 75t-75 181q0 27 18.5 45.5t45 18.5t45.5-18.5t19-45.5q0-53 37.5-90.5T512 256t90.5 37.5T640 384t-37.5 90.5T512 512q-26 0-45 39.5T448 640q0 27 18.5 45.5T512 704t45.5-18.5T576 640v-8q84-22 138-91t54-157q0-106-75-181t-181-75z"/>`),
		g.Group(children),
	)
}