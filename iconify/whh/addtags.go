package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Addtags(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1006.06 681l-325 324q-18 19-44.5 19t-44.5-19l-144-144v-61q0-40-28-68t-68-28h-32v-32q0-40-28-68t-68-28h-62l-144-144q-18-18-18-80V64q0-26 19-45t45-19h288q62 0 80 19l574 573q18 18 18 44.5t-18 44.5zm-750-553q-53 0-90.5 37.5t-37.5 90.5t37.5 90.5t90.5 37.5t90.5-37.5t37.5-90.5t-37.5-90.5t-90.5-37.5zm-192 640h64v-64q0-27 18.5-45.5t45-18.5t45.5 18.5t19 45.5v64h64q26 0 45 18.5t19 45.5t-19 45.5t-45 18.5h-64v64q0 26-19 45t-45.5 19t-45-19t-18.5-45v-64h-64q-27 0-45.5-18.5T.06 832t18.5-45.5t45.5-18.5z"/>`),
		g.Group(children),
	)
}