package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chef(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M958 446q-15 16-31 34t-40.5 48t-39.5 60t-15 52q0 27-93.5 45.5T512 704t-226.5-18.5T192 640q0-29-27-72.5t-50.5-70T68 448Q0 382 0 288q0-93 65.5-158.5T224 64q23 0 48 5q-16 41-16 91q0 107 66 178q-2-54-2-114q0-93 56-158.5T512 0t136 65.5T704 224q0 61-2 114q66-71 66-178q0-50-16-91q25-5 48-5q93 0 158.5 65.5T1024 288t-66 158zM511.5 768Q644 768 738 749t94-45v256q0 26-94 45t-226.5 19t-226-19t-93.5-45V704q0 27 93.5 45.5t226 18.5z"/>`),
		g.Group(children),
	)
}