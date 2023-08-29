package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicalNotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="m21.563.17l-11.11 3.399c-1.346.384-2.437 1.788-2.437 3.134v11.719s-.805-.543-2.598-.289C2.785 18.507.65 20.528.65 22.648s2.135 3.419 4.768 3.045c2.635-.372 4.566-2.331 4.566-4.452V11.235c0-.94 1.13-1.343 1.13-1.343l9.823-3.079s1.087-.365 1.087.641v8.037s-1.001-.576-2.794-.358c-2.633.319-4.768 2.298-4.768 4.417c0 2.121 2.135 3.463 4.768 3.143c2.635-.319 4.77-2.297 4.77-4.418V1.912C24 .566 22.908-.214 21.563.17z"/>`),
		g.Group(children),
	)
}