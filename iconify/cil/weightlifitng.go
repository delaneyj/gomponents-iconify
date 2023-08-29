package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weightlifitng(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="256" cy="156" r="36" fill="currentColor"/><path fill="currentColor" d="M400 16v56H112V16H16v144h96v-56h25.278l7.524 90.289l79.2 39.6V336h-79.138l-7.413 133.426H169.5L175.136 368h161.728l5.636 101.426h32.05L367.136 336H288V233.889l79.2-39.6L374.722 104H400v56h96V16ZM80 72v56H48V48h32Zm256.8 101.71l-80.8 40.4l-80.8-40.4l-5.811-69.71h173.222ZM464 128h-32V48h32Z"/>`),
		g.Group(children),
	)
}