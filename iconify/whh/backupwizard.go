package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backupwizard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1004.504 1005q-19 19-45 20t-44-17l-193-194q-18-18-17-44t20-45t45-20t44 17l193 194q18 18 17 44t-20 45zm-428-748v228l-193-194l33-33l-89-12l-25-53h210q27 0 45.5 18.5t18.5 45.5zm-320-64h18l-24 53l-58 8q1-26 19.5-43.5t44.5-17.5zm0 384q-26 0-45-19t-19-45V290l34 33l-17 94l73-43l202 203h-228zm384-384q0-27-18.5-46t-45.5-19h-384q-26 0-45 19t-19 46v384q0 26 19 45t45 19h356l128 128h-612q-26 0-45-19t-19-45V64q0-26 19-45t45-19h640q27 0 45.5 19t18.5 45v613l-128-128V193z"/>`),
		g.Group(children),
	)
}