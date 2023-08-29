package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Event(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.5 960h-896q-27 0-45.5-18.5T.5 896V384h1024v512q0 27-18.5 45.5T960.5 960zm-192-384h-256v192h256V576zM.5 192q0-27 19-45.5t45-18.5h64V64q0-26 19-45T193 0t45 19t18.5 45v64h512V64q0-26 19-45T833 0t45 19t18.5 45v64h64q27 0 45.5 19t18.5 45v128H.5V192z"/>`),
		g.Group(children),
	)
}