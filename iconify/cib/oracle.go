package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oracle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21.88 5.88H10.125C4.536 5.88.005 10.407 0 15.995c-.005 5.589 4.527 10.12 10.115 10.125H21.88C27.469 26.12 32 21.588 32 16S27.469 5.88 21.88 5.88zm-.255 16.672H10.38c-8.563-.172-8.563-12.932 0-13.104h11.245c8.735 0 8.735 13.104 0 13.104z"/>`),
		g.Group(children),
	)
}