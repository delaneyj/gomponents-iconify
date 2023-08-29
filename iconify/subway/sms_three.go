package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmsThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 0H64C28.6 0 0 28.6 0 64v256c0 35.4 28.6 64 64 64h128l-42.7 128l192-128H448c35.4 0 64-28.6 64-64V64c0-35.4-28.6-64-64-64zM128 234.7c-23.6 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7s42.7 19.1 42.7 42.7s-19.1 42.7-42.7 42.7zm128 0c-23.6 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7s42.7 19.1 42.7 42.7s-19.1 42.7-42.7 42.7zm128 0c-23.6 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7s42.7 19.1 42.7 42.7s-19.1 42.7-42.7 42.7z"/>`),
		g.Group(children),
	)
}