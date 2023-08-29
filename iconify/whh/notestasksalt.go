package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notestasksalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 1024H128q-53 0-90.5-37.5T0 896V256q0-53 37.5-90.5T128 128h64V64q0-27 18.5-45.5T256 0t45.5 18.5T320 64v64h128V64q0-27 18.5-45.5T512 0t45.5 18.5T576 64v64h128V64q0-27 18.5-45.5T768 0t45.5 18.5T832 64v64h64q53 0 90.5 37.5T1024 256v640q0 53-37.5 90.5T896 1024zm-82.5-622.5Q795 383 769.5 383T725 401L448 679L298 529q-19-18-44.5-18t-44 18.5t-18.5 44t19 43.5l183 183q6 11 8 14q19 18 47 17q27 1 46-17q2-3 8-14l312-311q18-18 18-43.5t-18.5-44z"/>`),
		g.Group(children),
	)
}