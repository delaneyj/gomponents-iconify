package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crontosignswiss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.87 10.1h4.72v4.72H9.87zm7.48 0h4.72v4.72h-4.72zm7.45 15.08h4.72v4.72H24.8zm7.48 0H37v4.72h-4.72zm-22.41 0h4.72v4.72H9.87zm7.48 0h4.72v4.72h-4.72zm7.45 7.48h4.72v4.72H24.8zm-14.93 0h4.72v4.72H9.87zm7.48 0h4.72v4.72h-4.72zm14.93 0H37v4.72h-4.72zM9.87 17.64h4.72v4.72H9.87zm7.48 0h4.72v4.72h-4.72zm11.71-7.54v4.26H24.8v3.68h4.26v4.26h3.67v-4.26h4.26v-3.68h-4.26V10.1h-3.67z"/>`),
		g.Group(children),
	)
}