package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terraform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m158.692 271.29l141.528 80.865V512l-141.528-80.867V271.29zm0-17.337L300.22 334.82V171.54L158.692 90.673v163.28zm155.577-82.662v164.862l142.937-82.143V89.713L314.269 171.29zM144.643 82.646L0 0v165.366l144.643 82.646V82.646z"/>`),
		g.Group(children),
	)
}