package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M192 48c0-26.5 21.5-48 48-48h160c26.5 0 48 21.5 48 48v464h-80v-80c0-26.5-21.5-48-48-48s-48 21.5-48 48v80h-80V48zM48 96h112v416H48c-26.5 0-48-21.5-48-48V320h80c8.8 0 16-7.2 16-16s-7.2-16-16-16H0v-64h80c8.8 0 16-7.2 16-16s-7.2-16-16-16H0v-48c0-26.5 21.5-48 48-48zm544 0c26.5 0 48 21.5 48 48v48h-80c-8.8 0-16 7.2-16 16s7.2 16 16 16h80v64h-80c-8.8 0-16 7.2-16 16s7.2 16 16 16h80v144c0 26.5-21.5 48-48 48H480V96h112zM312 64c-8.8 0-16 7.2-16 16v24h-24c-8.8 0-16 7.2-16 16v16c0 8.8 7.2 16 16 16h24v24c0 8.8 7.2 16 16 16h16c8.8 0 16-7.2 16-16v-24h24c8.8 0 16-7.2 16-16v-16c0-8.8-7.2-16-16-16h-24V80c0-8.8-7.2-16-16-16h-16z"/>`),
		g.Group(children),
	)
}