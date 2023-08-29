package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vinyl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm0-704q-80 0-136 56.5t-56 136T376 648t135.5 56t136-56T704 512.5t-56-136T512 320zm0 320q-53 0-90.5-37.5T384 512t37.5-90.5T512 384t90.5 37.5T640 512t-37.5 90.5T512 640zm-.5-192q-26.5 0-45 19T448 512.5t18.5 45t45 18.5t45.5-18.5t19-45t-19-45.5t-45.5-19z"/>`),
		g.Group(children),
	)
}