package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Miniwink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 576q0 94-69 167.5t-185.5 113T512 896t-257.5-39.5T69 743.5T0 576q0-27 19-45.5T64.5 512t45 18.5T128 576q0 47 51 91t141 72.5T512 768q103 0 192.5-27.5T845 669t51-93q0-27 19-45.5t45.5-18.5t45 18.5T1024 576zM832 256H640q-26 0-45-19t-19-45t19-45t45-19h192q27 0 45.5 19t18.5 45.5t-18.5 45T832 256zm-576 64q-53 0-90.5-47T128 160t37.5-113T256 0t90.5 47T384 160t-37.5 113t-90.5 47z"/>`),
		g.Group(children),
	)
}