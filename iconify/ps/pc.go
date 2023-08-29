package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M261 0H69Q42 0 23.5 18.5T5 64v384q0 27 18.5 45.5T69 512h192q28 0 46-18.5t18-45.5V64q0-27-18-45.5T261 0zm22 448q0 21-22 21H69q-21 0-21-21V213h235v235zm0-256H48V64q0-21 21-21h192q22 0 22 21v128zm-22 64q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15zM69 85h192v43H69V85zm171 75q0 11-11 11q-10 0-10-11t10-11q11 0 11 11zM91 427h64v21H91v-21zm85 0h64v21h-64v-21z"/>`),
		g.Group(children),
	)
}