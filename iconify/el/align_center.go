package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M290.625 99.202v178.006h618.75V99.202h-618.75zm-164.063 274.53v178.006h946.875V373.732H126.562zm107.813 274.53v178.006h731.25V648.262h-731.25zM0 922.792v178.006h1200V922.792H0z"/>`),
		g.Group(children),
	)
}