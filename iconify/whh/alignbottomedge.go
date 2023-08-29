package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alignbottomedge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.553 1024h-896q-26 0-45-19t-19-45.5t19-45t45-18.5h896q26 0 45 18.5t19 45t-19 45.5t-45 19zm-64-256h-256q-27 0-45.5-19t-18.5-45V320q0-27 18.5-45.5t45.5-18.5h256q26 0 45 18.5t19 45.5v384q0 27-19 45.5t-45 18.5zm-64-352q0-13-9.5-22.5t-22.5-9.5h-64q-13 0-22.5 9.5t-9.5 22.5v192q0 13 9.5 22.5t22.5 9.5h64q13 0 22.5-9.5t9.5-22.5V416zm-448 352h-256q-27 0-45.5-19t-18.5-45V64q0-26 18.5-45t45.5-19h256q26 0 45 19t19 45v640q0 27-19 45.5t-45 18.5z"/>`),
		g.Group(children),
	)
}