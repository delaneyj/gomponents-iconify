package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fax(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 5v5h-2V8H4v18h2v1c0 1.645 1.355 3 3 3s3-1.355 3-3v-1h16V10h-4V5H12zm2 2h8v5h-8V7zm-8 3h2v14H6V10zm4 2h2v2h12v-2h2v12H10V12zm3 4v2h2v-2h-2zm4 0v2h2v-2h-2zm4 0v2h2v-2h-2zm-8 4v2h2v-2h-2zm4 0v2h2v-2h-2zm4 0v2h2v-2h-2zM8 26h2v1c0 .555-.445 1-1 1c-.555 0-1-.445-1-1v-1z"/>`),
		g.Group(children),
	)
}