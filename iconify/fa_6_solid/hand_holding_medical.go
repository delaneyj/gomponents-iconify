package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandHoldingMedical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M224 24v56h-56c-13.3 0-24 10.7-24 24v48c0 13.3 10.7 24 24 24h56v56c0 13.3 10.7 24 24 24h48c13.3 0 24-10.7 24-24v-56h56c13.3 0 24-10.7 24-24v-48c0-13.3-10.7-24-24-24h-56V24c0-13.3-10.7-24-24-24h-48c-13.3 0-24 10.7-24 24zm335.7 368.2c17.8-13.1 21.6-38.1 8.5-55.9s-38.1-21.6-55.9-8.5L392.6 416H272c-8.8 0-16-7.2-16-16s7.2-16 16-16h80c17.7 0 32-14.3 32-32s-14.3-32-32-32H193.7c-29.1 0-57.3 9.9-80 28l-44.9 36H32c-17.7 0-32 14.3-32 32v64c0 17.7 14.3 32 32 32h320.5c29 0 57.3-9.3 80.7-26.5l126.6-93.3zm-367-8.2h.9h-.9z"/>`),
		g.Group(children),
	)
}