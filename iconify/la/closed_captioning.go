package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaptioning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 6v17.906h28V6zm2 2h24v13.906H4zm8 2c-2.75 0-5 2.25-5 5s2.25 5 5 5a4.973 4.973 0 0 0 3.125-1.125l-1.25-1.563c-.527.426-1.168.688-1.875.688c-1.668 0-3-1.332-3-3s1.332-3 3-3c.707 0 1.348.262 1.875.688l1.25-1.563A4.973 4.973 0 0 0 12 10zm10 0c-2.75 0-5 2.25-5 5s2.25 5 5 5a4.973 4.973 0 0 0 3.125-1.125l-1.25-1.563c-.527.426-1.168.688-1.875.688c-1.668 0-3-1.332-3-3s1.332-3 3-3c.707 0 1.348.262 1.875.688l1.25-1.563A4.973 4.973 0 0 0 22 10z"/>`),
		g.Group(children),
	)
}