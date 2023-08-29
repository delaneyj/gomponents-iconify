package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bullhorn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M480 32c0-12.9-7.8-24.6-19.8-29.6S434.5.2 425.3 9.3L381.7 53c-48 48-113.1 75-181 75H64c-35.3 0-64 28.7-64 64v96c0 35.3 28.7 64 64 64v128c0 17.7 14.3 32 32 32h64c17.7 0 32-14.3 32-32V352h8.7c67.9 0 133 27 181 75l43.6 43.6c9.2 9.2 22.9 11.9 34.9 6.9s19.8-16.6 19.8-29.6V300.4c18.6-8.8 32-32.5 32-60.4s-13.4-51.6-32-60.4V32zm-64 76.7v262.6C357.2 317.8 280.5 288 200.7 288H192v-96h8.7c79.8 0 156.5-29.8 215.3-83.3z"/>`),
		g.Group(children),
	)
}