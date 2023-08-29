package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volumedown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m794 761l-45-45q84-85 84-204.5T749 308l45-45q49 48 76 112.5T897 512t-27 136t-76 113zm-229 251L320 768V255L565 12q30-30 76 15v970q-45 45-76 15zM256 768H64q-26 0-45-18.5T0 704V319q0-26 19-45t45-19h192v513z"/>`),
		g.Group(children),
	)
}