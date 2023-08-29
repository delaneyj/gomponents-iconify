package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pytest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M47.289 47.818V0h80.312v47.818H47.29zm305.986 0V0h-80.313v47.818h80.313zm-112.837 0V0h-80.312v47.818h80.312zm249.044 15.287H22.58c-30.107 0-30.107 45.09 0 45.09h466.902c30.024 0 30.024-45.09 0-45.09zm-23.37-15.287V0h-80.313v47.818h80.312zm-80.313 75.664v117.134h80.312V123.482H385.8zm-112.837 0v201.852h80.313V123.482h-80.313zm-112.836 0v311.215h80.312V123.482h-80.312zm-112.837 0V512h80.312V123.482H47.29z"/>`),
		g.Group(children),
	)
}