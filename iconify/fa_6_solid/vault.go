package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vault(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 0C28.7 0 0 28.7 0 64v352c0 35.3 28.7 64 64 64h16l16 32h64l16-32h224l16 32h64l16-32h16c35.3 0 64-28.7 64-64V64c0-35.3-28.7-64-64-64H64zm160 320a80 80 0 1 0 0-160a80 80 0 1 0 0 160zm0-240a160 160 0 1 1 0 320a160 160 0 1 1 0-320zm256 141.3V336c0 8.8-7.2 16-16 16s-16-7.2-16-16V221.3c-18.6-6.6-32-24.4-32-45.3c0-26.5 21.5-48 48-48s48 21.5 48 48c0 20.9-13.4 38.7-32 45.3z"/>`),
		g.Group(children),
	)
}