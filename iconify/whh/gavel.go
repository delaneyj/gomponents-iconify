package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gavel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M959.998 320h-512v128h32q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-384q-13 0-22.5-9.5t-9.5-22.5t9.5-22.5t22.5-9.5h32V64h-32q-13 0-22.5-9.5t-9.5-22.5t9.5-22.5t22.5-9.5h384q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-32v128h512q26 0 45 18.5t19 45.5t-19 45.5t-45 18.5zm-704-224q0-13-9.5-22.5t-22.5-9.5t-22.5 9.5t-9.5 22.5v64q0 13 9.5 22.5t22.5 9.5t22.5-9.5t9.5-22.5V96zm0 256q0-13-9.5-22.5t-22.5-9.5t-22.5 9.5t-9.5 22.5v64q0 13 9.5 22.5t22.5 9.5t22.5-9.5t9.5-22.5v-64zm256 352q0 22 10 39.5t22 26t22 17t10 13.5v64q0 13-9.5 22.5t-22.5 9.5h-512q-13 0-22.5-9.5t-9.5-22.5v-64q0-5 10-13.5t22-17t22-26t10-39.5h448z"/>`),
		g.Group(children),
	)
}