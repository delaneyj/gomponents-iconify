package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Poundalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1023q-104 0-199-40.5t-163.5-109T40.5 710T0 511t40.5-198.5t109-163T313 40.5T512 0t199 40.5t163.5 109t109 163T1024 511t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm-64-448h64q26 0 45-18.5t19-45t-19-45.5t-45-19H400l-2-6q-14-28-14-58q0-53 37.5-90.5T512 255t90.5 37.5T640 383q0 27 18.5 45.5t45 18.5t45.5-18.5t19-45.5q0-106-75-181t-181-75t-181 75t-75 181q0 29 9 64h-9q-26 0-45 19t-19 45.5t19 45t45 18.5h64q0 35-10 66t-22 49t-22 39t-10 38q0 27 19 45.5t45 18.5h384q26 0 45-18.5t19-45t-19-45.5t-45-19H412q16-15 24.5-47t10.5-57z"/>`),
		g.Group(children),
	)
}