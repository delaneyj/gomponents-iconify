package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VersionPatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18.586 18.414L16 15.834l-2.588 2.582L12 17l4-4l4 4l-1.414 1.414z"/><path fill="currentColor" d="M1.586 17.414C1.196 17.024 1 16.512 1 16s.195-1.024.586-1.414l13-13C14.976 1.196 15.488 1 16 1s1.024.195 1.414.586l13 13c.39.39.586.902.586 1.414s-.195 1.024-.586 1.414l-13 13c-.39.39-.902.586-1.414.586s-1.024-.195-1.414-.586l-13-13ZM16 3L3 16l13 13l13-13L16 3Z"/>`),
		g.Group(children),
	)
}