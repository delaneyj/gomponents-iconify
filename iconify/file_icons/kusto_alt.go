package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KustoAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 71.01v378.358l286.526 46.916V15.716L0 71.01zm167.559 274.238l-55.853-84.94v84.94H69.258V168.753h42.448v81.173l55.853-81.173h46.916l-61.438 86.246l73.447 90.25h-58.925zm137.29-263.087v120.467h30.683v191.854h-30.684v40.251H512V82.161H304.848zm95.421 312.32h-49.855V258.774h49.855v135.709zm67.598 0h-52.407V144.74h52.407v249.742z"/>`),
		g.Group(children),
	)
}