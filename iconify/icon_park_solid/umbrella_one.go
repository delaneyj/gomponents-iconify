package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUmbrellaOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M44 24c0-11.046-8.954-20-20-20S4 12.954 4 24h40Z"/><path stroke-linecap="round" d="M24 24v14.554c0 3.014 2.486 5.457 5.5 5.457s5.5-2.443 5.5-5.457"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUmbrellaOne0)"/>`),
		g.Group(children),
	)
}