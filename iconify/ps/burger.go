package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Burger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M432 237v-51q21-11 21-37q0-62-43.5-105.5T304 0H176Q114 0 70.5 43.5T27 149q0 26 21 37v49q-18 3-30.5 17.5T5 286q0 20 14 35t33 18q-4 18-4 24q0 9 4 25q-25 12-25 39v21q0 27 18 45.5T91 512h298q28 0 46-18.5t18-45.5v-21q0-27-25-39q4-16 4-25q0-8-4-22q19-3 33-18t14-35q0-19-12.5-33.5T432 237zM91 192h128v43H91v-43zm170 0h128v43H261v-43zM176 43h128q45 0 76 31t31 75H69q0-44 31-75t76-31zm213 320q0 21-21 21H112q-21 0-21-21q0-22 21-22h256q21 0 21 22zm22 85q0 21-22 21H91q-22 0-22-21v-21h342v21zm10-149H59q-11 0-11-11t11-11h362q11 0 11 11t-11 11zM155 64h21v21h-21V64zm64 0h21v21h-21V64zm64 0h21v21h-21V64z"/>`),
		g.Group(children),
	)
}