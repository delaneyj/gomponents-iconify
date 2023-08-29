package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Doclets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M284.591 3.086H0v462.303c28.42-178.657 121.8-234.532 262.467-244.083C278.935 408.85 156.385 491.001 26.412 508.914h258.179c298.67-32.18 307.72-467.614 0-505.828z"/>`),
		g.Group(children),
	)
}