package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M964 756q43 63 56.5 124.5T1012 983t-78 41H92q-56 0-78-41T5.5 880.5T62 756l259-328V128h-64q-26 0-45-18.5t-19-45T212 19t45-19h512q26 0 45 19t19 45.5t-19 45t-45 18.5h-64v300zM577 448V128H449v320L206 760q-32 33-44.5 63.5T165 875t64 21h568q47 0 63.5-21t4-51.5T819 760z"/>`),
		g.Group(children),
	)
}