package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Removefolderalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.11 896h-896q-26 0-45-19t-19-45V256h1024v576q0 26-19 45t-45 19zm-82-407q18-18 18-43.5t-18-43.5t-43.5-18t-43.5 18l-87 87l-87-87q-18-18-43.5-18t-43.5 18t-18 43.5t18 43.5l87 87l-87 87q-18 18-18 43.5t18 43.5t43.5 18t43.5-18l87-87l87 87q18 18 43.5 18t43.5-18t18-43.5t-18-43.5l-87-87zM.11 64q0-27 19-45.5t45-18.5h384q27 0 45.5 18.5t18.5 45.5t18.5 45.5t45.5 18.5h384q26 0 45 19t19 45H.11V64z"/>`),
		g.Group(children),
	)
}