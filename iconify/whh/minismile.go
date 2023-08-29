package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minismile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 576q0 94-69 167.5t-185.5 113T512 896t-257.5-39.5T69 743.5T0 576q0-27 18.5-45.5t45-18.5t45.5 18.5t19 45.5q0 47 51 91t141 72.5T512 768q103 0 192.5-27.5T845 669t51-93q0-27 18.5-45.5t45-18.5t45.5 18.5t19 45.5zM768 320q-53 0-90.5-47T640 160t37.5-113T768 0t90.5 47T896 160t-37.5 113t-90.5 47zm-512 0q-53 0-90.5-47T128 160t37.5-113T256 0t90.5 47T384 160t-37.5 113t-90.5 47z"/>`),
		g.Group(children),
	)
}