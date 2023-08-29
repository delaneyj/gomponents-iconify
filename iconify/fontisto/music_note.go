package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.07.169a.913.913 0 0 0-.848-.114l.006-.002L6.151 3.991a.924.924 0 0 0-.613.869V17.7a3.847 3.847 0 0 0-1.846-.471h-.001a3.553 3.553 0 0 0-3.692 3.376v.008a3.555 3.555 0 0 0 3.699 3.385h-.007a3.553 3.553 0 0 0 3.692-3.376V7.731l9.23-3.223v8.973a3.86 3.86 0 0 0-1.846-.47h-.001a3.551 3.551 0 0 0-3.691 3.376v.008a3.554 3.554 0 0 0 3.699 3.385h-.007l.105.002a3.622 3.622 0 0 0 3.513-2.74l.005-.025a.908.908 0 0 0 .069-.34V.923a.921.921 0 0 0-.388-.752l-.003-.002z"/>`),
		g.Group(children),
	)
}