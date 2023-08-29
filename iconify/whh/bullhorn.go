package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bullhorn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M949 1011L657 768H256v192q0 27-19 45.5t-45 18.5h-64q-27 0-45.5-18.5T64 960V768q-27 0-45.5-18.5T0 704V320q0-27 18.5-45.5T64 256h593L949 13Q960 1 978.5.5t32 7.5t13.5 20v969q0 12-13.5 20t-32 7t-29.5-13zM480 256H352q13 0 22.5 9.5T384 288v448q0 13-9.5 22.5T352 768h128q-14 0-23-9.5t-9-22.5V288q0-13 9-22.5t23-9.5z"/>`),
		g.Group(children),
	)
}