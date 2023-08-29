package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 80c-8.8 0-16 7.2-16 16v162c5.1-1.3 10.5-2 16-2h384c5.5 0 10.9.7 16 2V96c0-8.8-7.2-16-16-16H64zM48 320v96c0 8.8 7.2 16 16 16h384c8.8 0 16-7.2 16-16v-96c0-8.8-7.2-16-16-16H64c-8.8 0-16 7.2-16 16zm-48 0V96c0-35.3 28.7-64 64-64h384c35.3 0 64 28.7 64 64v320c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64v-96zm280 48a24 24 0 1 1 48 0a24 24 0 1 1-48 0zm120-24a24 24 0 1 1 0 48a24 24 0 1 1 0-48z"/>`),
		g.Group(children),
	)
}