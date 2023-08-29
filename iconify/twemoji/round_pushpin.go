package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundPushpin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<ellipse cx="18" cy="34.5" fill="#292F33" rx="4" ry="1.5"/><path fill="#99AAB5" d="M14.339 10.725S16.894 34.998 18.001 35c1.106.001 3.66-24.275 3.66-24.275h-7.322z"/><circle cx="18" cy="8" r="8" fill="#DD2E44"/>`),
		g.Group(children),
	)
}