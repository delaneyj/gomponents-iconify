package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cpu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M151 22v32h18V22h-18zm48 0v32h18V22h-18zm48 0v32h18V22h-18zm48 0v32h18V22h-18zm48 0v32h18V22h-18zM131.7 73L73 131.7v248.6l58.7 58.7h248.6l58.7-58.7V131.7L380.3 73H131.7zm60.3 46h128v18H192v-18zM22 151v18h32v-18H22zm436 0v18h32v-18h-32zm-339 32h274v146H119V183zm-97 16v18h32v-18H22zm436 0v18h32v-18h-32zm-321 2v110h238V201H137zM22 247v18h32v-18H22zm436 0v18h32v-18h-32zM22 295v18h32v-18H22zm436 0v18h32v-18h-32zM22 343v18h32v-18H22zm436 0v18h32v-18h-32zm-266 32h128v18H192v-18zm-41 83v32h18v-32h-18zm48 0v32h18v-32h-18zm48 0v32h18v-32h-18zm48 0v32h18v-32h-18zm48 0v32h18v-32h-18z"/>`),
		g.Group(children),
	)
}