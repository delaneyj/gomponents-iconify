package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 0h178.7c17 0 33.3 6.7 45.3 18.7L365.3 96c12 12 18.7 28.3 18.7 45.3V448c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V64C0 28.7 28.7 0 64 0zm32 192c-17.7 0-32 14.3-32 32v32h64v-64H96zM64 352h256v-64H64v64zm256-128c0-17.7-14.3-32-32-32h-32v64h64v-32zm-160-32v64h64v-64h-64zm128 256c17.7 0 32-14.3 32-32v-32h-64v64h32zm-128-64v64h64v-64h-64zm-96 32c0 17.7 14.3 32 32 32h32v-64H64v32z"/>`),
		g.Group(children),
	)
}