package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonBiking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M400 96a48 48 0 1 0 0-96a48 48 0 1 0 0 96zm27.2 64l-61.8-48.8c-17.3-13.6-41.7-13.8-59.1-.3l-83.1 64.2c-30.7 23.8-28.5 70.8 4.3 91.6l60.5 38.4V416c0 17.7 14.3 32 32 32s32-14.3 32-32V288c0-10.7-5.3-20.7-14.2-26.6L295 232.9l60.3-48.5L396 217c5.7 4.5 12.7 7 20 7h64c17.7 0 32-14.3 32-32s-14.3-32-32-32h-52.8zM56 384a72 72 0 1 1 144 0a72 72 0 1 1-144 0zm200 0a128 128 0 1 0-256 0a128 128 0 1 0 256 0zm184 0a72 72 0 1 1 144 0a72 72 0 1 1-144 0zm200 0a128 128 0 1 0-256 0a128 128 0 1 0 256 0z"/>`),
		g.Group(children),
	)
}