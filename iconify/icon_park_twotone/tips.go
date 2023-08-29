package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tips(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTips0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M40 20c0 6.808-4.252 12.622-10.244 14.934H18.244C12.252 32.622 8 26.808 8 20c0-8.837 7.163-16 16-16s16 7.163 16 16Z"/><path d="m29.756 34.934l-.68 8.15a1 1 0 0 1-.996.916h-8.16a1 1 0 0 1-.996-.917l-.68-8.15M18 17v6l6-3l6 3v-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTips0)"/>`),
		g.Group(children),
	)
}