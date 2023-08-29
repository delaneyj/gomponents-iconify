package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fax(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M128 64v96h64V64h194.7L416 93.3V160h64V93.3c0-17-6.7-33.3-18.7-45.3L432 18.7C420 6.7 403.7 0 386.7 0H192c-35.3 0-64 28.7-64 64zM0 160v320c0 17.7 14.3 32 32 32h32c17.7 0 32-14.3 32-32V160c0-17.7-14.3-32-32-32H32c-17.7 0-32 14.3-32 32zm480 32H128v288c0 17.7 14.3 32 32 32h320c17.7 0 32-14.3 32-32V224c0-17.7-14.3-32-32-32zm-224 64a32 32 0 1 1 0 64a32 32 0 1 1 0-64zm96 32a32 32 0 1 1 64 0a32 32 0 1 1-64 0zm32 96a32 32 0 1 1 0 64a32 32 0 1 1 0-64zm-160 32a32 32 0 1 1 64 0a32 32 0 1 1-64 0z"/>`),
		g.Group(children),
	)
}