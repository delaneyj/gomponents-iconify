package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckVerified(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCheckVerified0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCheckVerified1" fill="currentColor"><path id="feCheckVerified2" d="M4.252 14H4a2 2 0 1 1 0-4h.252c.189-.734.48-1.427.856-2.064l-.18-.179a2 2 0 1 1 2.83-2.828l.178.179A7.952 7.952 0 0 1 10 4.252V4a2 2 0 1 1 4 0v.252c.734.189 1.427.48 2.064.856l.179-.18a2 2 0 1 1 2.828 2.83l-.179.178c.377.637.667 1.33.856 2.064H20a2 2 0 1 1 0 4h-.252a7.952 7.952 0 0 1-.856 2.064l.18.179a2 2 0 1 1-2.83 2.828l-.178-.179a7.952 7.952 0 0 1-2.064.856V20a2 2 0 1 1-4 0v-.252a7.952 7.952 0 0 1-2.064-.856l-.179.18a2 2 0 1 1-2.828-2.83l.179-.178A7.952 7.952 0 0 1 4.252 14ZM9 10l-2 2l4 4l6-6l-2-2l-4 4l-2-2Z"/></g></g>`),
		g.Group(children),
	)
}