package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picasa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M503 271L231 24C271 9 313 0 358 0c51 0 101 12 145 31v240zM329 163L12 453c-8-30-12-61-12-94C0 221 78 100 193 41zm213 314V51c104 62 175 177 175 308c0 41-8 81-20 118H542zM171 360v304C105 624 55 564 26 492zm38 155h472c-58 119-181 202-323 202c-53 0-104-13-149-33V515z"/>`),
		g.Group(children),
	)
}