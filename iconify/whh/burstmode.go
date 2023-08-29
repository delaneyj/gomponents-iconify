package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Burstmode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992.268 768h-704q-13 0-22.5-9.5t-9.5-22.5v-32h704V256h32q13 0 22.5 9.5t9.5 22.5v448q0 13-9.5 22.5t-22.5 9.5zm-128-128h-704q-13 0-22.5-9.5t-9.5-22.5v-32h704V128h32q13 0 22.5 9.5t9.5 22.5v448q0 13-9.5 22.5t-22.5 9.5zm-128-128h-704q-13 0-22.5-9.5T.268 480V32q0-13 9.5-22.5t22.5-9.5h704q13 0 22.5 9.5t9.5 22.5v448q0 13-9.5 22.5t-22.5 9.5zm-32-448h-640v384h640V64z"/>`),
		g.Group(children),
	)
}