package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosCalendarOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" d="M80 112v64h352v-64"/><path d="M352 96V64h-16v32H176V64h-16v32H64v352h384V96h-96zm80 336H80V192h352v240zm0-256H80v-64h80v32h16v-32h160v32h16v-32h80v64z" fill="currentColor"/>`),
		g.Group(children),
	)
}