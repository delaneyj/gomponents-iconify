package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Molecule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M800 1024q-93 0-158.5-65.5T576 800q0-97 72-163l-59-137q-37 12-77 12t-77-12l-59 137q72 67 72 163q0 93-65.5 158.5T224 1024T65.5 958.5T0 800t65.5-158.5T224 576q23 0 48 5l61-142q-77-75-77-183q0-106 75-181T512 0t181 75t75 181q0 107-77 183l61 142q25-5 48-5q93 0 158.5 65.5T1024 800t-65.5 158.5T800 1024zM191.5 640q-26.5 0-45 19T128 704t18.5 45t45 19t45.5-18.5t19-45.5t-19-45.5t-45.5-18.5zm256-512q-26.5 0-45 18.5T384 192t18.5 45.5t45 18.5t45.5-18.5t19-45.5t-19-45.5t-45.5-18.5zm320 512q-26.5 0-45 18.5t-18.5 45t18.5 45.5t45 19t45.5-18.5t19-45.5t-19-45.5t-45.5-18.5z"/>`),
		g.Group(children),
	)
}