package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GavelLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M23.7 10.41a1 1 0 0 1-.71-.29l-7.43-7.43A1 1 0 0 1 17 1.28l7.44 7.43a1 1 0 0 1-.71 1.7ZM11.86 22.25a1 1 0 0 0-.29-.71l-7.43-7.43a1 1 0 0 0-1.42 1.42L10.15 23a1 1 0 0 0 1.42 0a1 1 0 0 0 .29-.75ZM21.93 34H3a1 1 0 0 1-1-1.27l1.13-4a1 1 0 0 1 1-.73H20.8a1 1 0 0 1 1 .73l1.13 4a1 1 0 0 1-.17.87a1 1 0 0 1-.83.4ZM4.31 32H20.6l-.6-2H4.87Zm28.8-4.56l-14-14l2.36-2.36l-6.95-6.95l-8.94 8.94L12.51 20l2.35-2.34l14 14a3 3 0 0 0 4.24 0a3 3 0 0 0 .01-4.22ZM8.4 13.07L14.52 7l4.11 4.11l-6.12 6.11Zm23.29 17.2a1 1 0 0 1-1.41 0l-14-14l1.41-1.41l14 14a1 1 0 0 1 0 1.41Z"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}