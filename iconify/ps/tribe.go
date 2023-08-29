package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tribe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M419 189q-4 0-12 2l-24-45q7-7 7-20q0-14-9-23t-23-9q-17 0-25 13L218 70v-5q0-13-9.5-22.5T186 33t-22.5 9.5T154 65v8l-72 31Q69 83 45 83q-18 0-30.5 12.5T2 126t12.5 30.5T45 169q9 0 21-5l71 87q-7 16-7 30q0 29 20.5 49.5T200 351q25 0 44-16.5t24-40.5l113-41q14 22 38 22q18 0 30.5-12.5T462 232t-12.5-30.5T419 189zM87 133q0-1 .5-3t.5-3l74-40q9 11 24 11q11 0 22-9l19 6q-33 2-56.5 26.5T147 180q0 8 4 24zm72 92v-1h1zm96 14q-21-27-55-27q-8 0-20 2q-10-15-10-34q0-26 18.5-44t44.5-18t44.5 18.5T296 180q0 20-11.5 36T255 239zm124-21l-104 37q43-25 43-75q0-40-30-65l38 13q1 13 10 21.5t22 8.5q5 0 8-1l24 43q-8 8-11 18z"/>`),
		g.Group(children),
	)
}