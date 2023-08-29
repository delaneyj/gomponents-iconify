package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondsCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 512h256q27 0 45.5-18.5T384 448V64q0-27-18.5-45.5T320 0H64Q37 0 18.5 18.5T0 64v384q0 27 18.5 45.5T64 512zM43 64q0-21 21-21h256q21 0 21 21v384q0 21-21 21H64q-21 0-21-21V64zm213 192l-64-85l-64 85l64 85z"/>`),
		g.Group(children),
	)
}