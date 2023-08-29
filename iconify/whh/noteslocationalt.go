package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Noteslocationalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 1024H128q-53 0-90.5-37.5T0 896V256q0-53 37.5-90.5T128 128h64V64q0-27 18.5-45.5T256 0t45.5 18.5T320 64v64h128V64q0-27 18.5-45.5T512 0t45.5 18.5T576 64v64h128V64q0-27 18.5-45.5T768 0t45.5 18.5T832 64v64h64q53 0 90.5 37.5T1024 256v640q0 53-37.5 90.5T896 1024zM512 320q-80 0-136 56t-56 136q0 37 25.5 94T404 712t64.5 84.5T512 832t43.5-35.5T620 712t58.5-106t25.5-94q0-80-56-136t-136-56zM384 512q0-53 37.5-90.5T512 384t90.5 37.5T640 512t-37.5 90.5T512 640t-90.5-37.5T384 512z"/>`),
		g.Group(children),
	)
}