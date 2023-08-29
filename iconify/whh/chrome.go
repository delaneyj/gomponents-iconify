package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chrome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m512 1024l224-384q32-55 32-128q0-56-23-106t-64-86h305q38 93 38 192q0 139-68.5 257T769 955.5T512 1024zM261 461L109 197q71-91 176.5-144T512 0q140 0 257.5 69.5T955 256H512q-92 0-162.5 58.5T261 461zm251-141q80 0 136 56t56 136t-56 136t-136 56t-136-56t-56-136t56-136t136-56zM288 640q34 58 94 93t130 35q42 0 82-14l-153 265q-123-18-224-88.5t-159-181T0 512q0-137 69-256z"/>`),
		g.Group(children),
	)
}