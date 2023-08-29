package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Identi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M2 212q-1 85 60.5 144.5T212 416q50 0 92-21l158 47l-51-165q11-33 11-66q0-85-61.5-145T212 6T63.5 66.5T2 212zm352-1q0 27-9 49l27 75l2 6l9 25l-33-9l-72-22q-30 15-66 15q-59 0-101-40.5T69 211q0-57 42-98t101-41t100.5 40.5T354 211z"/>`),
		g.Group(children),
	)
}