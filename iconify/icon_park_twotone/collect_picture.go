package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectPicture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCollectPicture0"><g fill="none"><path fill="#fff" d="M44 21a2 2 0 1 0-4 0h4ZM23 8a2 2 0 1 0 0-4v4Zm16 32H9v4h30v-4ZM8 39V9H4v30h4Zm32-18v18h4V21h-4ZM9 8h14V4H9v4Zm0 32a1 1 0 0 1-1-1H4a5 5 0 0 0 5 5v-4Zm30 4a5 5 0 0 0 5-5h-4a1 1 0 0 1-1 1v4ZM8 9a1 1 0 0 1 1-1V4a5 5 0 0 0-5 5h4Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m6 35l10.693-9.802a2 2 0 0 1 2.653-.044L32 36m-4-5l4.773-4.773a2 2 0 0 1 2.615-.186L42 31"/><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M33.3 6C31.478 6 30 7.435 30 9.205c0 3.204 3.9 6.117 6 6.795c2.1-.678 6-3.59 6-6.795C42 7.435 40.523 6 38.7 6A3.326 3.326 0 0 0 36 7.362A3.326 3.326 0 0 0 33.3 6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCollectPicture0)"/>`),
		g.Group(children),
	)
}