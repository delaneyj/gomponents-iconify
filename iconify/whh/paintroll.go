package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paintroll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992.193 448h-480v192q26 0 45 19t19 45v256q0 26-19 45t-45 19h-64q-27 0-45.5-19t-18.5-45V704q0-26 18.5-45t45.5-19V416q0-13 9.5-22.5t22.5-9.5h480V192h-64v64q0 26-19 45t-45 19h-768q-27 0-45.5-18.5T.193 256V64q0-26 18.5-45t45.5-19h768q26 0 45 19t19 45v64h96q13 0 22.5 9.5t9.5 22.5v256q0 13-9.5 22.5t-22.5 9.5z"/>`),
		g.Group(children),
	)
}