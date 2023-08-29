package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Networksignal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m897 674l-96-84q95-109 95-253T801 85l96-85q60 69 93.5 155.5T1024 337t-33.5 181.5T897 674zM753 548l-97-85q48-55 48-126t-48-126l97-84q79 90 79 210t-79 211zm-241-83q-53 0-90.5-37.5T384 337t37.5-90.5T512 209t90.5 37.5T640 337t-37.5 90.5T512 465zm-241 83q-79-91-79-211t79-210l97 84q-48 55-48 126t48 126zm-48 42l-96 84q-60-69-93.5-155.5T0 337t33.5-181.5T127 0l96 85q-95 108-95 252t95 253z"/>`),
		g.Group(children),
	)
}