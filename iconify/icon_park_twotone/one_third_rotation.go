package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneThirdRotation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTOneThirdRotation0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M36 40.015A19.989 19.989 0 0 1 24 44c-7.403 0-13.866-4.022-17.324-10H15M32 5.664C39.064 8.75 44 15.8 44 24c0 3.643-.974 7.058-2.676 10l-4.042-7M4.099 26A20.239 20.239 0 0 1 4 24C4 12.954 12.954 4 24 4l-4.042 7"/><path fill="#555" d="M24 30a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTOneThirdRotation0)"/>`),
		g.Group(children),
	)
}